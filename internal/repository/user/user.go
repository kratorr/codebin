package user

import "database/sql"

type User struct {
	ID       int64
	Email    string
	Password string
	Username string
	Avatar   sql.NullString
	GoogleID sql.NullString
	GithubID sql.NullString
	TwitchID string
}
