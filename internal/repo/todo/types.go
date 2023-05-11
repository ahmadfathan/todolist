package todo

import (
	"database/sql"
)

type todoTable struct {
	TodoID          int64          `db:"todo_id"`
	ActivityGroupID int64          `db:"activity_group_id"`
	Title           string         `db:"title"`
	IsActive        bool           `db:"is_active"`
	Priority        string         `db:"priority"`
	CreatedAt       string         `db:"created_at"`
	UpdatedAt       sql.NullString `db:"updated_at"`
}
