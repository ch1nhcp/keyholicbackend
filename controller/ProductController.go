package controller

import (
	"encoding/json"
	"finalbackend/database"
	"finalbackend/models"
	"finalbackend/repository"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type People struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func Test(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var people People
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &people)
	json.NewEncoder(writer).Encode(people)
	fmt.Println(people)
}
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	var token string
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &token)
	json.NewEncoder(writer).Encode(token)
}

func GetAllProductPaginate(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	pages := params["page"]
	var value int
	if len(pages) == 0 {
		value = 1
	} else {
		page := strings.Split(pages[0], ",")
		// var arrConvert []int
		result, _ := strconv.Atoi(page[0])
		value = result
	}
	product := repository.GetAllProduct(value)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(product)
}

func AddNewProduct(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var product models.Product
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &product)
	err := repository.CheckProductExist(product.Name)
	if err != nil {
		json.NewEncoder(writer).Encode("product name exists")
		return
	}
	database.DB.Create(&product)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(product)
}
func GetProductByName(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	name := vars["name"]

	data, err := repository.FindProductByName(name)
	if err != nil {
		json.NewEncoder(writer).Encode("can't find")
		return
	}
	// writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(data)
}
func GetProductByCategory(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	pages := params["page"]
	var value int
	if len(pages) == 0 {
		value = 1
	} else {
		page := strings.Split(pages[0], ",")
		// var arrConvert []int
		result, _ := strconv.Atoi(page[0])
		value = result
	}
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	name := vars["name"]
	data := repository.FindProductByCategory(name, value)
	json.NewEncoder(writer).Encode(data)
}
func GetProductByManyBrand(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	pages := params["page"]
	var value int
	if len(pages) == 0 {
		value = 1
	} else {
		page := strings.Split(pages[0], ",")
		// var arrConvert []int
		result, _ := strconv.Atoi(page[0])
		value = result
	}
	writer.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	name := vars["name"]
	a := strings.Split(name, ",")
	// newname2 := "'" + newname1 + "'"

	data := repository.FindProductByBrand(a, value)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(data)
}
func GetProductLatest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data := repository.GetProductLatest()
	json.NewEncoder(writer).Encode(data)
}
func GetProductHot(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	data := repository.GetProductHot()
	json.NewEncoder(writer).Encode(data)
}
func GetProductSearch(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := request.URL.Query()
	search := params["name"]
	key := "%" + search[0] + "%"
	// page := strings.Split(search[0], ",")
	data := repository.GetProductSearch(key)
	json.NewEncoder(writer).Encode(data)
}
