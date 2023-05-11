package utils

import (
	"encoding/json"
	"errors"
	"html"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/ahmadfathan/todolist/internal/entity"
	"github.com/go-chi/chi"
)

func ParseActivityIDParam(r *http.Request) (int64, error) {

	var activityID int64

	activityIDStr := chi.URLParam(r, "activity-id")
	if activityIDStr == "" {
		return activityID, errors.New(ErrRequiredParameterIsEmpty)
	}

	activityID, err := strconv.ParseInt(activityIDStr, 10, 64)

	if err != nil {
		return activityID, errors.New(ErrInvalidValue)
	}

	return activityID, nil
}

func ParseCreateActivityRequest(r *http.Request) (entity.CreateActivityRequest, error) {

	var createActivityRequest entity.CreateActivityRequest

	unmarshalJSONBody(r, &createActivityRequest)

	if createActivityRequest.Title == "" {
		return createActivityRequest, entity.ErrTitleCanotBeNull
	}

	return createActivityRequest, nil
}

func ParseUpdateActivityRequest(r *http.Request) (int64, entity.UpdateActivityRequest, error) {

	var activityID int64
	var updateActivityRequest entity.UpdateActivityRequest

	unmarshalJSONBody(r, &updateActivityRequest)

	if updateActivityRequest.Title == "" {
		return activityID, updateActivityRequest, entity.ErrTitleCanotBeNull
	}

	activityIDStr := chi.URLParam(r, "activity-id")
	if activityIDStr == "" {
		return activityID, updateActivityRequest, errors.New(ErrRequiredParameterIsEmpty)
	}

	activityID, err := strconv.ParseInt(activityIDStr, 10, 64)

	if err != nil {
		return activityID, updateActivityRequest, errors.New(ErrInvalidValue)
	}

	return activityID, updateActivityRequest, nil
}

func ParseGetAllTodoRequest(r *http.Request) (int64, error) {

	activityGroupIDStr := getFormData(r, "activity_group_id")

	activityGroupID, err := strconv.ParseInt(activityGroupIDStr, 10, 64)

	if err != nil {
		return activityGroupID, errors.New(ErrInvalidValue)
	}

	return activityGroupID, nil
}

func ParseTodoIDParam(r *http.Request) (int64, error) {

	var todoID int64

	todoIDStr := chi.URLParam(r, "todo-id")
	if todoIDStr == "" {
		return todoID, errors.New(ErrRequiredParameterIsEmpty)
	}

	todoID, err := strconv.ParseInt(todoIDStr, 10, 64)

	if err != nil {
		return todoID, errors.New(ErrInvalidValue)
	}

	return todoID, nil
}

func ParseCreateTodoRequest(r *http.Request) (entity.CreateTodoRequest, error) {

	var createTodoRequest entity.CreateTodoRequest

	unmarshalJSONBody(r, &createTodoRequest)

	if createTodoRequest.Title == "" {
		return createTodoRequest, entity.ErrTitleCanotBeNull
	}

	if createTodoRequest.ActivityGroupID == 0 {
		return createTodoRequest, entity.ErrActivityGroupIDCanotBeNull
	}

	return createTodoRequest, nil
}

func ParseUpdateTodoRequest(r *http.Request) (int64, entity.UpdateTodoRequest, error) {

	var todoID int64
	var updateTodoRequest entity.UpdateTodoRequest

	unmarshalJSONBody(r, &updateTodoRequest)

	todoIDStr := chi.URLParam(r, "todo-id")
	if todoIDStr == "" {
		return todoID, updateTodoRequest, errors.New(ErrRequiredParameterIsEmpty)
	}

	todoID, err := strconv.ParseInt(todoIDStr, 10, 64)

	if err != nil {
		return todoID, updateTodoRequest, errors.New(ErrInvalidValue)
	}

	return todoID, updateTodoRequest, nil
}

// unexported methods

func getFormData(r *http.Request, param string) string {
	formValue := r.FormValue(param)
	formValueTrimmed := strings.TrimSpace(formValue)
	return html.EscapeString(formValueTrimmed)
}

func unmarshalJSONBody(r *http.Request, result interface{}) error {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}
