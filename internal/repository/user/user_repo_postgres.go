package user

import (
	"codebin/internal/model"
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var ErrAlreadyExists = errors.New("user already exists")

type Repo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (username, password, email, activation_status) VALUES ($1, $2, $3, $4)`
	row := r.db.QueryRowxContext(ctx, query, user.Username, user.Password, user.Email, user.ActivationStatus)
	if row.Err() != nil {
		err := row.Err()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrAlreadyExists
			}
		}
		return errors.Wrap(row.Err(), "insert user into DB")
	}

	return nil
}

func (r *Repo) DeleteByID(ctx context.Context, i int) error {
	panic("implement me")
}

func (r *Repo) Update(ctx context.Context, user *model.User) error {
	panic("implement me")
	return nil
}

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return nil, nil
}

func (r *Repo) GetUserByName(ctx context.Context, name string) (user *model.User, err error) {
	log.Debug().Str("method", "GetUserByName").Str("name", name).Msg("get user from db")
	query := `SELECT id, username, password, email, activation_status FROM users WHERE username = $1`
	if err := r.db.GetContext(ctx, user, query, name); err != nil {
		return nil, errors.Wrap(err, "get user from DB")
	}
	return user, nil
}

func (r *Repo) UpdateActivationStatus(ctx context.Context, email string, status int) error {
	query := `UPDATE users SET activation_status = $1 WHERE email = $2`
	row := r.db.QueryRowxContext(ctx, query, status, email)
	if row.Err() != nil {
		return errors.Wrap(row.Err(), "insert user into DB")
	}
	return nil
}
