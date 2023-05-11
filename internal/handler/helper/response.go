package helper

import (
	"encoding/json"
	"net/http"
)

type GlobalResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, status string, message string, data interface{}) {

	response := &GlobalResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, status string, message string) {
	errResponse := &ErrorResponse{
		Status:  status,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errResponse)
}
