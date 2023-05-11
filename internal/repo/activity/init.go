package activity

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type activityRepo struct {
	db             *sqlx.DB
	readStatement  map[string]*sqlx.Stmt
	writeStatement map[string]*sqlx.Stmt
}

var (
	readStatementActivity = map[string]string{
		rStmtGetAllActivity:  queryGetAllActivity,
		rStmtGetActivityByID: queryGetActivityByID,
	}

	writeStatementActivity = map[string]string{
		wStmtInsertActivity: queryInsertActivity,
		wStmtUpdateActivity: queryUpdateActivity,
		wStmtDeleteActivity: queryDeleteActivity,
	}
)

func New(db *sqlx.DB) (*activityRepo, error) {
	var (
		ctx = context.Background()
	)

	readStmt := map[string]*sqlx.Stmt{}
	writeStmt := map[string]*sqlx.Stmt{}

	activityRepo := &activityRepo{
		db:             db,
		readStatement:  readStmt,
		writeStatement: writeStmt,
	}

	activityRepo.InitReadPrepareStatement(ctx)
	activityRepo.InitWritePrepareStatement(ctx)

	return activityRepo, nil
}
