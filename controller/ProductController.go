package controller

import (
	"encoding/json"
	"finalbackend/database"
	"finalbackend/models"
	"finalbackend/repository"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

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
	// writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(data)
}
func GetProductByBrand(writer http.ResponseWriter, request *http.Request) {
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
	newname1 := strings.ReplaceAll(name, ",", "','")
	newname2 := "'" + newname1 + "'"
	var brand models.Brand
	database.DB.Raw("SELECT id FROM `brands` WHERE name in (?)", name).Scan(&brand)
	data := repository.FindProductByBrand(newname2, value)
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(data)
	json.NewEncoder(writer).Encode(brand)
}