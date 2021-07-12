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

func AddNewImageproduct(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var imageproduct models.Imageproduct
	database.DB.Create(&imageproduct)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(imageproduct)
}
func GetImageproductById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	imageproduct, err := repository.GetImageproductById(id)
	if err != nil {
		json.NewEncoder(writer).Encode("not find")
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(imageproduct)
}

func GetAllImageproduct(writer http.ResponseWriter, request *http.Request) {
	imageproduct := repository.GetAllImageproduct()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(imageproduct)
}

func DelImageproductById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	repository.DelImageproductById(id)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("thanh cong")
}

func UpdateImageproduct(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var imageproduct models.Imageproduct
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &imageproduct)
	err := repository.UpdateImageproduct(&imageproduct)
	if err != nil {
		json.NewEncoder(writer).Encode("error")
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("thanh cong")
	}

}
