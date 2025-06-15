package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/surajmaity1/backend-api/dtos"
	"github.com/surajmaity1/backend-api/services"
	"net/http"
	"strconv"
)

func CreatePost(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody dtos.PostRequest

	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	var validate = validator.New()
	if err := validate.Struct(requestBody); err != nil {
		logrus.Error(err)
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(dtos.CustomErrorResponse(http.StatusBadRequest, err.Error()))
	}

	postId, err := services.CreatePost(request.Context(), requestBody)
	if err != nil {
		logrus.Error(err)
		json.NewEncoder(response).Encode(dtos.CustomErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	responseBody := dtos.PostResponse{
		Id:          postId,
		PostContent: requestBody.PostContent,
		PostImage:   requestBody.PostImage,
		PostedBy:    requestBody.PostedBy,
	}
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(responseBody)
}

func GetPost(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	response.Header().Add("Content-Type", "application/json")
	postId := params.ByName("id")
	id, err := strconv.ParseInt(postId, 10, 64)

	if err != nil {
		logrus.Error(err)
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(dtos.CustomErrorResponse(http.StatusBadRequest, "Invalid post id"))
		return
	}

	post, err := services.GetPost(request.Context(), id)

	if err != nil {
		logrus.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			response.WriteHeader(http.StatusNotFound)
			json.NewEncoder(response).Encode(dtos.CustomErrorResponse(http.StatusNotFound, "Post not found"))
			return
		}
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
