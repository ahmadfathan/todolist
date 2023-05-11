package activity

import "database/sql"

type activityTable struct {
	ActivityID int64          `db:"activity_id"`
	Title      string         `db:"title"`
	Email      string         `db:"email"`
	CreatedAt  string         `db:"created_at"`
	UpdatedAt  sql.NullString `db:"updated_at"`
}
