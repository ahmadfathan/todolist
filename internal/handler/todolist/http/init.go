package http

import (
	"context"

	"github.com/ahmadfathan/todolist/internal/entity"
)

type ActivityUseCase interface {
	GetAllActivity(ctx context.Context) ([]entity.Activity, error)
	GetActivityByID(ctx context.Context, activityID int64) (entity.Activity, error)
	CreateActivity(ctx context.Context, createActivityRequest entity.CreateActivityRequest) (entity.Activity, error)
	UpdateActivity(ctx context.Context, activityID int64, updateActivityRequest entity.UpdateActivityRequest) (entity.Activity, error)
	DeleteActivity(ctx context.Context, activityID int64) error
}

type TodoUseCase interface {
	GetAllTodo(ctx context.Context, activityGroupID int64) ([]entity.Todo, error)
	GetTodoByID(ctx context.Context, todoID int64) (entity.Todo, error)
	CreateTodo(ctx context.Context, createdTodoRequest entity.CreateTodoRequest) (entity.Todo, error)
	UpdateTodo(ctx context.Context, todoID int64, updateTodoRequest entity.UpdateTodoRequest) (entity.Todo, error)
	DeleteTodo(ctx context.Context, todoID int64) error
}

type Handler struct {
	activityUc ActivityUseCase
	todoUc     TodoUseCase
}

func New(activityUc ActivityUseCase, todoUc TodoUseCase) *Handler {
	return &Handler{
		activityUc: activityUc,
		todoUc:     todoUc,
	}
}
