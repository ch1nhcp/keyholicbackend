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
	"time"

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
	cookie, _ := request.Cookie("jwt")
	json.NewEncoder(writer).Encode(cookie)
}
func SetCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "1",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode("ok")
}
func DelCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	json.NewEncoder(w).Encode("log out successfully")
}
func GetAllProduct(writer http.ResponseWriter, request *http.Request) {
	product := repository.GetProduct()
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(product)
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
