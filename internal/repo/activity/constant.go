package activity

// Error
const (
	// db error
	ErrPrepareStatement = "db: error prepare statement"
)

// DB query
const (
	// read query
	queryGetAllActivity  = "SELECT * FROM activities"
	queryGetActivityByID = `SELECT * FROM activities 
							WHERE activity_id = ?;`

	// write query
	queryInsertActivity = `	INSERT INTO activities (title, email)
							VALUES (?,?);`

	queryUpdateActivity = `	UPDATE activities
							SET title=COALESCE(NULLIF(?, ''), title), email=COALESCE(NULLIF(?, ''), email), updated_at=?
							WHERE activity_id=?;`

	queryDeleteActivity = ` DELETE FROM activities
							WHERE activity_id=?;`
)

// DB statement
const (
	// read statement
	rStmtGetAllActivity  = "get_all_activity"
	rStmtGetActivityByID = "get_activity_by_id"

	// write statement
	wStmtInsertActivity = "insert_activity"
	wStmtUpdateActivity = "update_activity"
	wStmtDeleteActivity = "delete_activity"
)
