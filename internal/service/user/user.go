package user

import (
	"codebin/internal/model"
	userRepo "codebin/internal/repository/user"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

const (
	NotActivated = iota + 0
	Activated
	Deactivated
)

var ErrInternal = errors.New("create hash")

type UserRepo interface {
	Create(ctx context.Context, user *model.User) error
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	UpdateActivationStatus(ctx context.Context, email string, status int) error
}

type MailSender interface {
	SendEmail(ctx context.Context, receiver string, body, subject string) error
}

type Service struct {
	repo       UserRepo
	mailSender MailSender
	domainName string
}

func NewUserService(repo UserRepo, mailSender MailSender, domainName string) *Service {
	return &Service{
		repo:       repo,
		mailSender: mailSender,
		domainName: domainName,
	}
}

func (s *Service) Register(ctx context.Context, userReq *RegisterReq) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Err(err).Msg("create hash")
		return ErrInternal
	}

	u, err := s.repo.GetUserByEmail(ctx, userReq.Email)
	if u != nil && err == nil {
		return ErrAlreadyExists
	}
	user := &model.User{
		Email:            userReq.Email,
		Username:         userReq.Username,
		Password:         string(hash),
		ActivationStatus: NotActivated,
	}

	if err = s.repo.Create(ctx, user); err != nil {
		if errors.Is(err, userRepo.ErrAlreadyExists) {
			return ErrAlreadyExists
		}
		log.Err(err).Msg("create user")
		return ErrInternal
	}

	//
	//if err = s.sendActivationEmail(ctx, user); err != nil {
	//	return errors.Wrap(err, "send activation email")
	//}   //TODO FIX send email

	return nil
}

func (s *Service) Login(ctx context.Context, userName, password string) (*model.User, error) {
	user, err := s.repo.GetUserByName(ctx, userName)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	if user.ActivationStatus == NotActivated || user.ActivationStatus == Deactivated {
		return nil, ErrUserNotActivated
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("password is incorrect")
	}
	//TODO return refresh token
	return user, nil
}

func (s *Service) Activate(ctx context.Context, email string) error {
	if err := s.repo.UpdateActivationStatus(ctx, email, Activated); err != nil {
		return errors.Wrap(err, "update user")
	}

	return nil
}

func (s *Service) sendActivationEmail(ctx context.Context, user *model.User) error {
	activationToken, err := createActivationToken(user.Email)
	if err != nil {

	}
	body := fmt.Sprintf(emailTemplate, user.Username, "https://"+s.domainName+"/v1/auth/activate?token="+activationToken)
	if err = s.mailSender.SendEmail(ctx, user.Email, body, emailSubject); err != nil {
	}

	return nil
}

func createActivationToken(userEmail string) (string, error) {
	claims := jwt.MapClaims{
		"email": userEmail,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="))
}
