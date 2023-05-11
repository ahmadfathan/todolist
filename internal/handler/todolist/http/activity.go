package http

import (
	"fmt"
	"net/http"

	"github.com/ahmadfathan/todolist/internal/entity"
	"github.com/ahmadfathan/todolist/internal/handler/helper"
	"github.com/ahmadfathan/todolist/pkg/utils"
)

func (h *Handler) GetAllActivity(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	activities, _ := h.activityUc.GetAllActivity(ctx)

	helper.WriteSuccessResponse(w, http.StatusOK, "Success", "Success", activities)
}

func (h *Handler) GetActivityByID(w http.ResponseWriter, r *http.Request) {

	activityID, _ := utils.ParseActivityIDParam(r)

	ctx := r.Context()

	activity, err := h.activityUc.GetActivityByID(ctx, activityID)

	if err == entity.ErrNotFound {

		errorMsg := fmt.Sprintf("Activity with ID %d Not Found", activityID)
		helper.WriteErrorResponse(w, http.StatusNotFound, "Not Found", errorMsg)
		return
	}

	helper.WriteSuccessResponse(w, http.StatusOK, "Success", "Success", activity)
}

func (h *Handler) CreateActivity(w http.ResponseWriter, r *http.Request) {

	createActivityRequest, err := utils.ParseCreateActivityRequest(r)

	if err == entity.ErrTitleCanotBeNull {
		helper.WriteErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	ctx := r.Context()

	activity, _ := h.activityUc.CreateActivity(ctx, createActivityRequest)

	helper.WriteSuccessResponse(w, http.StatusCreated, "Success", "Success", activity)
}

func (h *Handler) UpdateActivity(w http.ResponseWriter, r *http.Request) {

	activityID, updateActivityRequest, err := utils.ParseUpdateActivityRequest(r)

	if err == entity.ErrTitleCanotBeNull {
		helper.WriteErrorResponse(w, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}

	ctx := r.Context()

	activity, err := h.activityUc.UpdateActivity(ctx, activityID, updateActivityRequest)

	if err == entity.ErrNotFound {

		errorMsg := fmt.Sprintf("Activity with ID %d Not Found", activityID)
		helper.WriteErrorResponse(w, http.StatusNotFound, "Not Found", errorMsg)
		return
	}

	helper.WriteSuccessResponse(w, http.StatusOK, "Success", "Success", activity)
}

func (h *Handler) DeleteActivity(w http.ResponseWriter, r *http.Request) {

	activityID, _ := utils.ParseActivityIDParam(r)

	ctx := r.Context()

	err := h.activityUc.DeleteActivity(ctx, activityID)

	if err == entity.ErrNotFound {

		errorMsg := fmt.Sprintf("Activity with ID %d Not Found", activityID)
		helper.WriteErrorResponse(w, http.StatusNotFound, "Not Found", errorMsg)
		return
	}

	helper.WriteSuccessResponse(w, http.StatusOK, "Success", "Success", map[string]string{})
}
