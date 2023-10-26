package app

import (
	swaggerUI "codebin/swagger_ui"
	"context"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"time"

	"codebin/config"
	"codebin/internal/interceptor"
	"codebin/internal/mailsender"
	snippetRepo "codebin/internal/repository/snippet"
	userRepo "codebin/internal/repository/user"
	usecaseSnippet "codebin/internal/service/snippet"
	usecaseUser "codebin/internal/service/user"
	g "codebin/internal/transport/grpc"
	"codebin/pkg/db/postgresql"
	pbAuth "codebin/pkg/proto/auth"
	pbSnipptet "codebin/pkg/proto/snippet"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	grpcHealth "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func Run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	fmt.Println(cfg)
	ctx := context.Background()
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	//
	// ssoGoogle := &oauth2.Config{
	//	ClientID:     cfg.Oauth2Google.ClientID,
	//	ClientSecret: cfg.Oauth2Google.ClientSecret,
	//	Endpoint:     google.Endpoint,
	//	RedirectURL:  cfg.Oauth2Google.RedicrectURI,
	//	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	//}
	time.Sleep(time.Second * 2)
	db, err := postgresql.NewPostgresClient(
		ctx,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
	)
	if err != nil {
		panic(err)
	}
	// mailSender := mailsender.NewSender()
	// mailSender.SendEmail("smirnovv91@bk.ru", "hello world")

	snippetRepo := snippetRepo.NewSnippetRepo(db)
	snippetService := usecaseSnippet.NewSnippetUseCase(snippetRepo)
	snippetImlp := g.NewSnippetHandler(&g.SnippetHandlerDeps{Service: snippetService})

	uRepo := userRepo.NewRepo(db)

	mailSenderDeps := &mailsender.SenderDeps{
		SenderEmail: cfg.MailSender.From,
		Password:    cfg.MailSender.Password,
		SMTPHost:    cfg.MailSender.SMTPHost,
		SMTPPort:    cfg.MailSender.SMTPPort,
	}
	mailSender := mailsender.NewSender(mailSenderDeps)
	userService := usecaseUser.NewUserService(uRepo, mailSender, cfg.Domain.Name)
	authImpl := g.NewAuthHandler(&g.AuthHandlerDeps{
		Service:               userService,
		RefreshTokenSecretKey: cfg.Secret.RefreshSecretKey,
		AccessTokenSecretKey:  cfg.Secret.AccessSecretKey,
	},
	)

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(interceptor.NewAuthInterceptor(cfg.Secret.AccessSecretKey))),
		grpc.Creds(insecure.NewCredentials()),
	}
	grpcServer := grpc.NewServer(opts...)

	reflection.Register(grpcServer)
	grpcHealth.RegisterHealthServer(grpcServer, health.NewServer())

	pbSnipptet.RegisterSnippetV1Server(grpcServer, snippetImlp)
	pbAuth.RegisterAuthServer(grpcServer, authImpl)

	rpcMux := runtime.NewServeMux()
	optsHTTP := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pbSnipptet.RegisterSnippetV1HandlerFromEndpoint(ctx, rpcMux, ":9000", optsHTTP); err != nil {
		log.Fatal(err)
	}
	if err := pbAuth.RegisterAuthHandlerFromEndpoint(ctx, rpcMux, ":9000", optsHTTP); err != nil {
		log.Fatal(err)
	}

	staticFs, err := fs.Sub(swaggerUI.SwaggerUI, ".")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", rpcMux)
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.FS(staticFs))))

	errGroup := errgroup.Group{}

	errGroup.Go(func() error {
		lis, err := net.Listen("tcp", ":9000")
		if err != nil {
			return err
		}

		err = grpcServer.Serve(lis)
		if err != nil {
			return err
		}

		return nil
	})

	srv := &http.Server{
		Handler:      mux,
		Addr:         ":9001",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	errGroup.Go(func() error {
		return srv.ListenAndServe()
	})

	err = errGroup.Wait()
	if err != nil {
		return err
	}
	return nil
}
