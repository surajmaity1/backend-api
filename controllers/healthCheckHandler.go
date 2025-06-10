package controllers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/surajmaity1/backend-api/utils"
	"net/http"
	"time"
)

type HealthCheckResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func HealthCheckHandler(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Set("Content-Type", "application/json")

	data := HealthCheckResponse{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if err := utils.WriteResponse(data, response); err != nil {
		http.Error(response, `{"status":"error","message":"Internal Server Error"}`, http.StatusInternalServerError)
		return
	}
}
