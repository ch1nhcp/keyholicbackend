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

func AddNewOrderitem(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var orderitem models.Orderitem
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &orderitem)
	database.DB.Create(&orderitem)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(orderitem)
}
func GetOrderitemById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	orderitem, err := repository.GetOrderitemById(id)
	if err != nil {
		json.NewEncoder(writer).Encode("not find")
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(orderitem)
}

func GetAllOrderitem(writer http.ResponseWriter, request *http.Request) {
	orderitem := repository.GetAllOrderitem()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(orderitem)
}

func DelOrderitemById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	repository.DelOrderitemById(id)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("thanh cong")
}

func UpdateOrderitem(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var orderitem models.Orderitem
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &orderitem)
	err := repository.UpdateOrderitem(&orderitem)
	if err != nil {
		json.NewEncoder(writer).Encode("error")
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("thanh cong")
	}

}
