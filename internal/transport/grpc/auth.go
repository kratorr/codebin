package grpc

import (
	"context"
	"time"

	"codebin/internal/model"
	"codebin/internal/service"
	"codebin/internal/service/user"
	"codebin/pkg/jwt"
	pbAuth "codebin/pkg/proto/auth"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 5 * time.Minute
)

type Auth interface {
	Register(ctx context.Context, user *user.RegisterReq) error
	Login(ctx context.Context, userName, password string) (*model.User, error)
	Activate(ctx context.Context, email string) error
}

type AuthHandler struct {
	pbAuth.UnimplementedAuthServer
	service               Auth
	refreshTokenSecretKey string
	accessTokenSecretKey  string
}

type AuthHandlerDeps struct {
	Service               Auth
	RefreshTokenSecretKey string
	AccessTokenSecretKey  string
}

func NewAuthHandler(deps *AuthHandlerDeps) *AuthHandler {
	return &AuthHandler{
		service:               deps.Service,
		accessTokenSecretKey:  deps.AccessTokenSecretKey,
		refreshTokenSecretKey: deps.RefreshTokenSecretKey,
	}
}

func (a *AuthHandler) Register(ctx context.Context, req *pbAuth.RegisterRequest) (*emptypb.Empty, error) {
	registerReq := &user.RegisterReq{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		Email:    req.GetEmail(),
	}

	err := a.service.Register(ctx, registerReq)
	if err != nil {
		if errors.Is(err, service.ErrAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}
		return nil, status.Error(codes.Internal, "register user")
	}
	return &emptypb.Empty{}, nil
}

func (a *AuthHandler) Login(ctx context.Context, req *pbAuth.LoginRequest) (*pbAuth.LoginResponse, error) {
	u, err := a.service.Login(ctx, req.GetUsername(), req.GetPassword())
	if err != nil {
		if errors.Is(err, user.ErrUserNotActivated) {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return nil, status.Error(codes.Internal, "unknown error")
	}
	log.Info().Msgf("user %s logged in", u.Username)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, status.Errorf(codes.Aborted, "wrong password")
		}
	}
	userModel := &model.User{
		Username: u.Username,
		ID:       u.ID,
		Email:    u.Email,
	}

	refreshToken, err := jwt.GenerateToken(userModel, []byte(a.refreshTokenSecretKey), refreshTokenExpiration)
	if err != nil {
		return nil, errors.Wrap(err, "generate refresh token")
	}
	accessToken, err := jwt.GenerateToken(userModel, []byte(a.accessTokenSecretKey), accessTokenExpiration)
	if err != nil {
		return nil, errors.Wrap(err, "generate access token")
	}
	return &pbAuth.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (a *AuthHandler) GetRefreshToken(
	_ context.Context,
	req *pbAuth.GetRefreshTokenRequest,
) (*pbAuth.GetRefreshTokenResponse, error) {
	claims, err := jwt.VerifyToken(req.GetRefreshToken(), []byte("secret"))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid refresh token")
	}

	refreshToken, err := jwt.GenerateToken(&model.User{
		Username: claims.Username,
		ID:       claims.UserID,
	},
		[]byte(a.refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return nil, err
	}

	return &pbAuth.GetRefreshTokenResponse{RefreshToken: refreshToken}, nil
}

func (a *AuthHandler) GetAccessToken(
	_ context.Context,
	req *pbAuth.GetAccessTokenRequest,
) (*pbAuth.GetAccessTokenResponse, error) {
	claims, err := jwt.VerifyToken(req.GetRefreshToken(), []byte(a.refreshTokenSecretKey))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid refresh token")
	}

	accessToken, err := jwt.GenerateToken(&model.User{
		Username: claims.Username,
		ID:       claims.UserID,
	},
		[]byte(a.accessTokenSecretKey),
		accessTokenExpiration,
	)
	if err != nil {
		return nil, err
	}

	return &pbAuth.GetAccessTokenResponse{AccessToken: accessToken}, nil
}

func (a *AuthHandler) Activate(ctx context.Context, req *pbAuth.ActiveRequest) (*emptypb.Empty, error) {
	claims, err := jwt.VerifyToken(req.GetToken(), []byte(a.accessTokenSecretKey))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid token")
	}
	if err = a.service.Activate(ctx, claims.Email); err != nil {
		return nil, status.Error(codes.Internal, "activate user")
	}

	return &emptypb.Empty{}, nil
}
