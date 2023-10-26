package interceptor

import (
	"context"
	"slices"
	"strings"

	"codebin/pkg/jwt"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const authPrefix = "Bearer "

var allowedMethods = []string{
	"/Auth/Login",
	"/Auth/Register",
	"/Auth/Activate",
	"/Auth/GetAccessToken",
	"/Auth/GetRefreshToken",
}

type key int

var userKey key

type User struct {
	ID       int64
	Username string
}

func NewAuthInterceptor(accessTokenSecretKey string) grpc.UnaryServerInterceptor {
	interceptor := func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Info().Msgf("grpc method: %s", info.FullMethod)
		if slices.Contains(allowedMethods, info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
		}

		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {
			return handler(ctx, req)
			//return nil, status.Errorf(codes.Unauthenticated, "missing authentication token")
		}
		if !strings.HasPrefix(authHeader[0], authPrefix) {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization header format")
		}

		accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

		claims, err := jwt.VerifyToken(accessToken, []byte(accessTokenSecretKey))
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "access token is invalid")
		}

		u := User{
			ID:       claims.UserID,
			Username: claims.Username,
		}
		ctx = context.WithValue(ctx, userKey, u)

		//ctx = context.WithValue(ctx, "userID", claims.UserID)
		//ctx = context.WithValue(ctx, "username", claims.Username)

		return handler(ctx, req)
	}

	return interceptor
}

func FromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userKey).(*User)
	return u, ok
}
