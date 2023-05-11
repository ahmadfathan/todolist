package activity

import (
	"context"

	"github.com/ahmadfathan/todolist/internal/entity"
)

type activityRepo interface {
	GetAllActivity(ctx context.Context) ([]entity.Activity, error)
	GetActivityByID(ctx context.Context, activityID int64) (entity.Activity, error)
	InsertActivity(ctx context.Context, createActivityRequest entity.CreateActivityRequest) (entity.Activity, error)
	UpdateActivity(ctx context.Context, activityID int64, updateActivityRequest entity.UpdateActivityRequest) (entity.Activity, error)
	DeleteActivity(ctx context.Context, activityID int64) error
}

type ActivityUseCase struct {
	activityRepo activityRepo
}

func New(activityRepo activityRepo) *ActivityUseCase {
	return &ActivityUseCase{
		activityRepo: activityRepo,
	}
}

func (u *ActivityUseCase) GetAllActivity(ctx context.Context) ([]entity.Activity, error) {
	return u.activityRepo.GetAllActivity(ctx)
}

func (u *ActivityUseCase) GetActivityByID(ctx context.Context, activityID int64) (entity.Activity, error) {
	return u.activityRepo.GetActivityByID(ctx, activityID)
}

func (u *ActivityUseCase) CreateActivity(ctx context.Context, createActivityRequest entity.CreateActivityRequest) (entity.Activity, error) {
	return u.activityRepo.InsertActivity(ctx, createActivityRequest)
}

func (u *ActivityUseCase) UpdateActivity(ctx context.Context, activityID int64, updateActivityRequest entity.UpdateActivityRequest) (entity.Activity, error) {
	_, err := u.activityRepo.GetActivityByID(ctx, activityID)

	if err == entity.ErrNotFound {
		return entity.Activity{}, err
	}

	return u.activityRepo.UpdateActivity(ctx, activityID, updateActivityRequest)
}

func (u *ActivityUseCase) DeleteActivity(ctx context.Context, activityID int64) error {

	_, err := u.activityRepo.GetActivityByID(ctx, activityID)

	if err == entity.ErrNotFound {
		return err
	}

	return u.activityRepo.DeleteActivity(ctx, activityID)
}
