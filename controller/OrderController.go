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

func AddNewOrder(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var order models.Order
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &order)
	database.DB.Create(&order)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(order)
}
func GetOrderById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	order, err := repository.GetOrderById(id)
	if err != nil {
		json.NewEncoder(writer).Encode("not find")
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(order)
}
func GetOrderByIdUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	order := repository.GetOrderByIdUser(id)

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(order)
}
func GetAllOrder(writer http.ResponseWriter, request *http.Request) {
	order := repository.GetAllOrder()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(order)
}

func DelOrderById(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	strIdProduct := vars["id"]
	id, _ := strconv.Atoi(strIdProduct)
	repository.DelOrderById(id)
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode("thanh cong")
}

func UpdateOrder(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var order models.Order
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &order)
	err := repository.UpdateOrder(&order)
	if err != nil {
		json.NewEncoder(writer).Encode("error")
	} else {
		writer.WriteHeader(http.StatusOK)
		json.NewEncoder(writer).Encode("thanh cong")
	}

}
