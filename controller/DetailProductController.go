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

func AddNewDetailproduct(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var detailproduct models.Detailproduct
	database.DB.Create(&detailproduct)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(detailproduct)
}
func GetDetailproductById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	detailproduct, err := repository.GetDetailproductById(id)
	if err != nil {
		json.NewEncoder(writer).Encode("not find")
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(detailproduct)
}

func GetAllDetailproduct(writer http.ResponseWriter, request *http.Request) {
	detailproduct := repository.GetAllDetailproduct()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(detailproduct)
}

func DelDetailproductById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	repository.DelDetailproductById(id)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("thanh cong")
}

func UpdateDetailproduct(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var detailproduct models.Detailproduct
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &detailproduct)
	err := repository.UpdateDetailproduct(&detailproduct)
	if err != nil {
		json.NewEncoder(writer).Encode("error")
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("thanh cong")
	}

}
