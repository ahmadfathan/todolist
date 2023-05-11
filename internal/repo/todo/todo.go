package todo

import (
	"context"
	"time"

	"github.com/ahmadfathan/todolist/internal/entity"
	"github.com/jmoiron/sqlx"
)

func (r *todoRepo) InitReadPrepareStatement(ctx context.Context) (devErrMsg string, err error) {
	readStmt := map[string]*sqlx.Stmt{}

	for k, v := range readStatementTodo {
		readStmt[k], err = r.db.PreparexContext(ctx, v)

		if err != nil {
			return ErrPrepareStatement, err
		}
	}

	r.readStatement = readStmt

	return "", nil
}

func (r *todoRepo) InitWritePrepareStatement(ctx context.Context) (devErrMsg string, err error) {
	writeStmt := map[string]*sqlx.Stmt{}

	for k, v := range writeStatementTodo {
		writeStmt[k], err = r.db.PreparexContext(ctx, v)
		if err != nil {
			return ErrPrepareStatement, err
		}
	}

	r.writeStatement = writeStmt

	return "", nil
}

func (r *todoRepo) GetAllTodo(ctx context.Context) ([]entity.Todo, error) {

	var (
		todos      []entity.Todo
		todoTables []todoTable
	)

	r.readStatement[rStmtGetAllTodo].SelectContext(ctx, &todoTables)

	todos = r.castTodoTablesToEntities(todoTables)

	return todos, nil
}

func (r *todoRepo) GetAllTodoByActivityGroupID(ctx context.Context, activityGroupID int64) ([]entity.Todo, error) {

	var (
		todos      []entity.Todo
		todoTables []todoTable
	)

	r.readStatement[rStmtGetAllTodoByActivityGroupID].SelectContext(ctx, &todoTables, activityGroupID)

	todos = r.castTodoTablesToEntities(todoTables)

	return todos, nil
}

func (r *todoRepo) GetTodoByID(ctx context.Context, todoID int64) (entity.Todo, error) {

	var (
		todo       entity.Todo
		todoTables []todoTable
	)

	r.readStatement[rStmtGetTodoByID].SelectContext(ctx, &todoTables, todoID)

	if len(todoTables) == 0 {
		return todo, entity.ErrNotFound
	}

	todo = r.castTodoTableToEntity(todoTables[0])

	return todo, nil
}

func (r *todoRepo) InsertTodo(ctx context.Context, createdTodoRequest entity.CreateTodoRequest) (entity.Todo, error) {

	var (
		todoID int64
		todo   entity.Todo
	)

	defaultIsActive := true
	defaultPriority := "very-high"

	result, _ := r.writeStatement[wStmtInsertTodo].ExecContext(ctx, createdTodoRequest.Title, createdTodoRequest.ActivityGroupID, defaultPriority, defaultIsActive)

	todoID, _ = result.LastInsertId()

	todo, _ = r.GetTodoByID(ctx, todoID)

	return todo, nil
}

func (r *todoRepo) UpdateTodo(ctx context.Context, todoID int64, updateTodoRequest entity.UpdateTodoRequest) (entity.Todo, error) {
	var (
		todo entity.Todo
	)

	currentTime := time.Now().Format("2006-01-02 15:04:05") // get current time

	r.writeStatement[wStmtUpdateTodo].ExecContext(
		ctx,
		updateTodoRequest.Title,
		updateTodoRequest.ActivityGroupID,
		updateTodoRequest.Priority,
		&updateTodoRequest.IsActive,
		currentTime,
		todoID,
	)

	todo, _ = r.GetTodoByID(ctx, todoID)

	return todo, nil
}
func (r *todoRepo) DeleteTodo(ctx context.Context, todoID int64) error {
	r.writeStatement[wStmtDeleteTodo].ExecContext(ctx, todoID)

	return nil
}

// unexported method

func (r *todoRepo) castTodoTableToEntity(todoTable todoTable) entity.Todo {
	var todo entity.Todo

	todo.TodoID = todoTable.TodoID
	todo.ActivityGroupID = todoTable.ActivityGroupID
	todo.Title = todoTable.Title
	todo.IsActive = todoTable.IsActive
	todo.Priority = todoTable.Priority
	todo.CreatedAt = todoTable.CreatedAt

	if todoTable.UpdatedAt.Valid {
		todo.UpdatedAt = todoTable.UpdatedAt.String
	}

	return todo
}

func (r *todoRepo) castTodoTablesToEntities(todoTables []todoTable) []entity.Todo {
	var todos []entity.Todo

	for _, v := range todoTables {
		todo := r.castTodoTableToEntity(v)
		todos = append(todos, todo)
	}

	return todos
}
