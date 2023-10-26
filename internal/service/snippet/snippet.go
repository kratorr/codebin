package usecase

import (
	"codebin/internal/model"
	"context"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const urlLength = 8

type SnippetRepo interface {
	Create(ctx context.Context, snippet *model.Snippet, user *model.User) error
	DeleteByID(ctx context.Context, id int) error
	ReadLatest(ctx context.Context) ([]*model.Snippet, error)
	ReadByUser(ctx context.Context, userID int) ([]*model.Snippet, error)
}

type ImageGenerator interface {
	CreateImage(code string)
}

type SnippetUseCase struct {
	repo SnippetRepo
	//imageGenerator ImageGenerator
}

func NewSnippetUseCase(repo SnippetRepo) *SnippetUseCase {
	return &SnippetUseCase{repo: repo}
}

func (s *SnippetUseCase) DeleteSnippetByID(ctx context.Context, id string) error {
	//TODO implement me
	//проверить есть ли пользователь, является ли пользователь владельцем сниппета
	panic("implement me")
}

func (s *SnippetUseCase) CreateSnippet(ctx context.Context, snippet *model.Snippet, model *model.User) (*model.Snippet, error) {
	snippet.ShortURL = s.generateRandomString(urlLength)
	log.Info().Msgf("create snippet with title: %s, url: %s", snippet.Title, snippet.ShortURL)
	if err := s.repo.Create(ctx, snippet, model); err != nil {
		return snippet, errors.Wrap(err, "create snippet in db")
	}
	return snippet, nil
}

func (s *SnippetUseCase) ReadLatest(ctx context.Context) ([]*model.Snippet, error) {
	resp, _ := s.repo.ReadLatest(ctx)
	return resp, nil
}

func (s *SnippetUseCase) generateRandomString(length int) string {
	rand.NewSource(time.Now().UnixNano())
	characters := []rune(charSet)
	result := make([]rune, length)
	for i := range result {
		result[i] = characters[rand.Intn(len(characters))]
	}

	return string(result)
}
