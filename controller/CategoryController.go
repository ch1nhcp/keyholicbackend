package controller

import (
	"encoding/json"
	"finalbackend/database"
	"finalbackend/models"
	"finalbackend/repository"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddNewCategory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var category models.Category
	database.DB.Create(&category)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(category)
}
func GetCategoryById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	category, err := repository.GetCategoryById(id)
	if err != nil {
		json.NewEncoder(writer).Encode("not find")
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(category)
}

func GetAllCategory(writer http.ResponseWriter, request *http.Request) {
	category := repository.GetAllCategory()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(category)
}

func DelCategoryById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	repository.DelCategoryById(id)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("thanh cong")
}

func UpdateCategory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var category models.Category
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &category)
	err := repository.UpdateCategory(&category)
	if err != nil {
		json.NewEncoder(writer).Encode("error")
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("thanh cong")
	}

}
