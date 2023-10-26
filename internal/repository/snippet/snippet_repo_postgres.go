package snippet

import (
	"codebin/internal/model"
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var ErrCreateSnippet = errors.New("failed to create snippet")

type Repo struct {
	db *sqlx.DB
}

func NewSnippetRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, snippet *model.Snippet, user *model.User) error {
	var userID interface{}
	if user != nil {
		userID = user.ID
	}
	query := `INSERT INTO snippets (title, text, language_id, user_id, url, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.ExecContext(ctx, query, snippet.Title, snippet.Text, 1, userID, snippet.ShortURL, time.Now())
	if err != nil {
		log.Err(err).Msg("insert snippet in db")
		return ErrCreateSnippet
	}
	return nil
}

func (r *Repo) DeleteByID(ctx context.Context, id int) error {
	//query := `DELETE FROM snippets WHERE id = $1`
	//rows, err := r.db.Query(ctx, query, id)
	//if err != nil {
	//	log.Err(err).Msg("delete snippet in db")
	//	return err
	//}
	//rows.Close()
	return nil
}

func (r *Repo) ReadLatest(ctx context.Context) ([]*model.Snippet, error) {
	query := `SELECT id, url FROM snippets LIMIT 10`

	result := make([]*Snippet, 0)
	if err := r.db.SelectContext(ctx, &result, query); err != nil {
		return nil, errors.Wrap(err, "select snippets latest from db")
	}
	snippets := make([]*model.Snippet, len(result))
	for v := range result {
		snippets[v] = &model.Snippet{
			ID: result[v].ID,
			//Title:     result[v].Title.
			Text: result[v].Text,
			//	Language:  model.Language{ID: result[v].LanguageID},
			ShortURL: result[v].ShortURL,
			//	CreatedAt: result[v].CreatedAt,
		}
	}
	fmt.Println(snippets[0].ShortURL)
	return snippets, nil
}

func (r *Repo) ReadByUser(ctx context.Context, userID int) ([]*model.Snippet, error) {
	//TODO implement me
	panic("implement me")
}
