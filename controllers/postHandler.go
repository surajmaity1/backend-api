package controllers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/surajmaity1/backend-api/dtos"
	"github.com/surajmaity1/backend-api/services"
	"net/http"
)

func CreatePost(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var requestBody dtos.PostRequest

	if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	var validate = validator.New()
	if err := validate.Struct(requestBody); err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	postId, err := services.CreatePost(request.Context(), requestBody)
	if err != nil {
		logrus.Error(err)
		http.Error(response, "Internal server error", http.StatusInternalServerError)
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

//func GetPost(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
//	body, err := io.ReadAll(request.Body)
//	fmt.Println(string(body))
//
//	if err != nil {
//		logrus.Error(err)
//		http.Error(response, "Internal server error", http.StatusInternalServerError)
//		return
//	}
//
//	services.GetPost(1)
//
//	response.Header().Add("Content-Type", "application/json")
//	response.WriteHeader(http.StatusOK)
//}
