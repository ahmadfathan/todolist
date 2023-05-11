package activity

import (
	"context"
	"time"

	"github.com/ahmadfathan/todolist/internal/entity"
	"github.com/jmoiron/sqlx"
)

func (r *activityRepo) InitReadPrepareStatement(ctx context.Context) (devErrMsg string, err error) {
	readStmt := map[string]*sqlx.Stmt{}

	for k, v := range readStatementActivity {
		readStmt[k], err = r.db.PreparexContext(ctx, v)

		if err != nil {
			return ErrPrepareStatement, err
		}
	}

	r.readStatement = readStmt

	return "", nil
}

func (r *activityRepo) InitWritePrepareStatement(ctx context.Context) (devErrMsg string, err error) {
	writeStmt := map[string]*sqlx.Stmt{}

	for k, v := range writeStatementActivity {
		writeStmt[k], err = r.db.PreparexContext(ctx, v)

		if err != nil {
			return ErrPrepareStatement, err
		}
	}

	r.writeStatement = writeStmt

	return "", nil
}

func (r *activityRepo) GetAllActivity(ctx context.Context) ([]entity.Activity, error) {

	var (
		activities     []entity.Activity
		activityTables []activityTable
	)

	r.readStatement[rStmtGetAllActivity].SelectContext(ctx, &activityTables)

	activities = r.castActivityTablesToEntities(activityTables)

	return activities, nil
}

func (r *activityRepo) GetActivityByID(ctx context.Context, activityID int64) (entity.Activity, error) {

	var (
		activity       entity.Activity
		activityTables []activityTable
	)

	r.readStatement[rStmtGetActivityByID].SelectContext(ctx, &activityTables, activityID)

	if len(activityTables) == 0 {
		return activity, entity.ErrNotFound
	}

	activity = r.castActivityTableToEntity(activityTables[0])

	return activity, nil
}

func (r *activityRepo) InsertActivity(ctx context.Context, createActivityRequest entity.CreateActivityRequest) (entity.Activity, error) {

	var (
		activityID int64
		activity   entity.Activity
	)

	result, _ := r.writeStatement[wStmtInsertActivity].ExecContext(ctx, createActivityRequest.Title, createActivityRequest.Email)

	activityID, _ = result.LastInsertId()

	activity, _ = r.GetActivityByID(ctx, activityID)

	return activity, nil
}

func (r *activityRepo) UpdateActivity(ctx context.Context, activityID int64, updateActivityRequest entity.UpdateActivityRequest) (entity.Activity, error) {

	var (
		activity entity.Activity
	)

	currentTime := time.Now().Format("2006-01-02 15:04:05") // get current time

	r.writeStatement[wStmtUpdateActivity].ExecContext(ctx, updateActivityRequest.Title, updateActivityRequest.Email, currentTime, activityID)

	activity, _ = r.GetActivityByID(ctx, activityID)

	return activity, nil
}

func (r *activityRepo) DeleteActivity(ctx context.Context, activityID int64) error {

	r.writeStatement[wStmtDeleteActivity].ExecContext(ctx, activityID)

	return nil
}

// unexported method

func (r *activityRepo) castActivityTableToEntity(activityTable activityTable) entity.Activity {
	var activity entity.Activity

	activity.ActivityID = activityTable.ActivityID
	activity.Title = activityTable.Title
	activity.Email = activityTable.Email
	activity.CreatedAt = activityTable.CreatedAt

	if activityTable.UpdatedAt.Valid {
		activity.UpdatedAt = activityTable.UpdatedAt.String
	}

	return activity
}

func (r *activityRepo) castActivityTablesToEntities(activityTables []activityTable) []entity.Activity {
	var activities []entity.Activity

	for _, v := range activityTables {
		activity := r.castActivityTableToEntity(v)
		activities = append(activities, activity)
	}

	return activities
}
