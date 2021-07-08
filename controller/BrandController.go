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

func AddNewBrand(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var brand models.Brand
	database.DB.Create(&brand)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(brand)
}
func GetBrandById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	brand, err := repository.GetBrandById(id)
	if err != nil {
		json.NewEncoder(writer).Encode("not find")
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(brand)
}

func GetAllBrand(writer http.ResponseWriter, request *http.Request) {
	brand := repository.GetAllBrand()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(brand)
}

func DelBrandById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	repository.DelBrandById(id)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("thanh cong")
}

func UpdateBrand(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var brand models.Brand
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &brand)
	err := repository.UpdateBrand(&brand)
	if err != nil {
		json.NewEncoder(writer).Encode("error")
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("thanh cong")
	}

}
