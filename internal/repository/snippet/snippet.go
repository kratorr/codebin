package snippet

import "database/sql"

type Snippet struct {
	ID        int64          `db:"id"`
	Text      string         `db:"text"`
	Language  string         `db:"language_id"`
	Title     sql.NullString `db:"title"`
	Author    sql.NullInt64  `db:"user_id"`
	ShortURL  string         `db:"url"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt sql.NullTime   `db:"updated_at"`
}
