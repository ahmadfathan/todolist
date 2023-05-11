package todo

// Error
const (
	// db error
	ErrPrepareStatement = "db: error prepare statement"
)

// DB query
const (
	// read query
	queryGetAllTodo                  = "SELECT * FROM todos"
	queryGetAllTodoByActivityGroupID = `SELECT * FROM todos 
										WHERE activity_group_id = ?;`
	queryGetTodoByID = `SELECT * FROM todos 
						WHERE todo_id = ?;`

	// write query
	queryInsertTodo = `	INSERT INTO todos (title, activity_group_id, priority, is_active)
						VALUES (?,?,?,?);`

	queryUpdateTodo = `	UPDATE todos
						SET 
							title=COALESCE(NULLIF(?, ''), title), 
							activity_group_id=COALESCE(NULLIF(?, 0), activity_group_id), 
							priority=COALESCE(NULLIF(?, ''), priority), 
							is_active=COALESCE(?, is_active), 
							updated_at=?
						WHERE todo_id=?;`

	queryDeleteTodo = ` DELETE FROM todos
						WHERE todo_id=?;`
)

// DB statement
const (
	// read statement
	rStmtGetAllTodo                  = "get_all_todo"
	rStmtGetAllTodoByActivityGroupID = "get_all_todo_by_activity_group_id"
	rStmtGetTodoByID                 = "get_todo_by_id"

	// write statement
	wStmtInsertTodo = "insert_todo"
	wStmtUpdateTodo = "update_todo"
	wStmtDeleteTodo = "delete_todo"
)
