package todo

import (
	"context"

	"github.com/ahmadfathan/todolist/internal/entity"
)

type todoRepo interface {
	GetAllTodo(ctx context.Context) ([]entity.Todo, error)
	GetAllTodoByActivityGroupID(ctx context.Context, activityGroupID int64) ([]entity.Todo, error)
	GetTodoByID(ctx context.Context, todoID int64) (entity.Todo, error)
	InsertTodo(ctx context.Context, createdTodoRequest entity.CreateTodoRequest) (entity.Todo, error)
	UpdateTodo(ctx context.Context, todoID int64, updateTodoRequest entity.UpdateTodoRequest) (entity.Todo, error)
	DeleteTodo(ctx context.Context, todoID int64) error
}

type TodoUseCase struct {
	todoRepo todoRepo
}

func New(todoRepo todoRepo) *TodoUseCase {
	return &TodoUseCase{
		todoRepo: todoRepo,
	}
}

func (u *TodoUseCase) GetAllTodo(ctx context.Context, activityGroupID int64) ([]entity.Todo, error) {
	if activityGroupID == 0 {
		return u.todoRepo.GetAllTodo(ctx)
	}

	return u.todoRepo.GetAllTodoByActivityGroupID(ctx, activityGroupID)
}

func (u *TodoUseCase) GetTodoByID(ctx context.Context, todoID int64) (entity.Todo, error) {
	return u.todoRepo.GetTodoByID(ctx, todoID)
}

func (u *TodoUseCase) CreateTodo(ctx context.Context, createdTodoRequest entity.CreateTodoRequest) (entity.Todo, error) {
	return u.todoRepo.InsertTodo(ctx, createdTodoRequest)
}

func (u *TodoUseCase) UpdateTodo(ctx context.Context, todoID int64, updateTodoRequest entity.UpdateTodoRequest) (entity.Todo, error) {
	_, err := u.todoRepo.GetTodoByID(ctx, todoID)

	if err == entity.ErrNotFound {
		return entity.Todo{}, err
	}

	return u.todoRepo.UpdateTodo(ctx, todoID, updateTodoRequest)
}

func (u *TodoUseCase) DeleteTodo(ctx context.Context, todoID int64) error {
	_, err := u.todoRepo.GetTodoByID(ctx, todoID)

	if err == entity.ErrNotFound {
		return err
	}

	return u.todoRepo.DeleteTodo(ctx, todoID)
}
