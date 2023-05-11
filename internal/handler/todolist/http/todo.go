package http

import (
	"fmt"
	"net/http"

	"github.com/ahmadfathan/todolist/internal/entity"
	"github.com/ahmadfathan/todolist/internal/handler/helper"
	"github.com/ahmadfathan/todolist/pkg/utils"
)

func (h *Handler) GetAllTodo(w http.ResponseWriter, r *http.Request) {

	activityGroupID, _ := utils.ParseGetAllTodoRequest(r)

	ctx := r.Context()

	todos, _ := h.todoUc.GetAllTodo(ctx, activityGroupID)

	if len(todos) == 0 {
		todos = make([]entity.Todo, 0)
	}

	helper.WriteSuccessResponse(w, http.StatusOK, "Success", "Success", todos)
}

func (h *Handler) GetTodoByID(w http.ResponseWriter, r *http.Request) {

	todoID, _ := utils.ParseTodoIDParam(r)

	ctx := r.Context()

	todo, err := h.todoUc.GetTodoByID(ctx, todoID)

	if err == entity.ErrNotFound {

		errorMsg := fmt.Sprintf("Todo with ID %d Not Found", todoID)
		helper.WriteErrorResponse(w, http.StatusNotFound, "Not Found", errorMsg)
		return
	}

	helper.WriteSuccessResponse(w, http.StatusOK, "Success", "Success", todo)
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	createTodoRequest, err := utils.ParseCreateTodoRequest(r)

	if err == entity.ErrTitleCanotBeNull {
		helper.WriteErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	if err == entity.ErrActivityGroupIDCanotBeNull {
		helper.WriteErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	ctx := r.Context()

	todo, _ := h.todoUc.CreateTodo(ctx, createTodoRequest)

	helper.WriteSuccessResponse(w, http.StatusCreated, "Success", "Success", todo)
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {

	todoID, updateTodoRequest, _ := utils.ParseUpdateTodoRequest(r)

	ctx := r.Context()

	todo, err := h.todoUc.UpdateTodo(ctx, todoID, updateTodoRequest)

	if err == entity.ErrNotFound {

		errorMsg := fmt.Sprintf("Todo with ID %d Not Found", todoID)
		helper.WriteErrorResponse(w, http.StatusNotFound, "Not Found", errorMsg)
		return
	}

	helper.WriteSuccessResponse(w, http.StatusOK, "Success", "Success", todo)
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {

	todoID, _ := utils.ParseTodoIDParam(r)

	ctx := r.Context()

	err := h.todoUc.DeleteTodo(ctx, todoID)

	if err == entity.ErrNotFound {

		errorMsg := fmt.Sprintf("Todo with ID %d Not Found", todoID)
		helper.WriteErrorResponse(w, http.StatusNotFound, "Not Found", errorMsg)
		return
	}

	helper.WriteSuccessResponse(w, http.StatusOK, "Success", "Success", map[string]string{})
}
