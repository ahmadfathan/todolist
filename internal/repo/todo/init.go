package todo

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type todoRepo struct {
	db             *sqlx.DB
	readStatement  map[string]*sqlx.Stmt
	writeStatement map[string]*sqlx.Stmt
}

var (
	readStatementTodo = map[string]string{
		rStmtGetAllTodo:                  queryGetAllTodo,
		rStmtGetAllTodoByActivityGroupID: queryGetAllTodoByActivityGroupID,
		rStmtGetTodoByID:                 queryGetTodoByID,
	}

	writeStatementTodo = map[string]string{
		wStmtInsertTodo: queryInsertTodo,
		wStmtUpdateTodo: queryUpdateTodo,
		wStmtDeleteTodo: queryDeleteTodo,
	}
)

func New(db *sqlx.DB) (*todoRepo, error) {
	var (
		ctx = context.Background()
	)

	readStmt := map[string]*sqlx.Stmt{}
	writeStmt := map[string]*sqlx.Stmt{}

	todoRepo := &todoRepo{
		db:             db,
		readStatement:  readStmt,
		writeStatement: writeStmt,
	}

	todoRepo.InitReadPrepareStatement(ctx)
	todoRepo.InitWritePrepareStatement(ctx)

	return todoRepo, nil
}
