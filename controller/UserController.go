package controller

import (
	"encoding/json"
	"finalbackend/models"
	"finalbackend/repository"
	"finalbackend/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user models.User
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &user)

	err := repository.RegisterUser(&user)
	fmt.Println(user)
	if err != nil {
		return
	} else {
		json.NewEncoder(writer).Encode("success")
	}
}

func Login(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var User models.User

	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &User)
	user, err := repository.Login(&User)
	if err != nil {
		json.NewEncoder(writer).Encode("sai tên đăng nhập hoặc mật khẩu")
		return
	}
	token, err := util.GenerateJwt(strconv.Itoa((user.Id)))
	if err != nil {
		panic(err.Error())
	}
	//đưa vào cookie
	//b1 : tọa cookie
	cookie := &http.Cookie{
		Name:    "token",
		Path:    "/",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
		// HttpOnly: true,
	}
	//bước 2 set cookie
	// writer.Header().Set("jwt", token)
	http.SetCookie(writer, cookie)
	writer.WriteHeader(http.StatusCreated)
<<<<<<< HEAD
	json.NewEncoder(writer).Encode("đăng nhập thành công")
	json.NewEncoder(writer).Encode(User)
=======
	// json.NewEncoder(writer).Encode("đăng nhập thành công")
	json.NewEncoder(writer).Encode(user)
>>>>>>> 28e18b7fe29baedf032232cbcd9cc754ca9c0425
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie := http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(writer, &cookie)
	json.NewEncoder(writer).Encode("log out successfully")
}

type cookie struct {
	Value string
}

func CheckCookie(writer http.ResponseWriter, request *http.Request) {
	var cookie cookie
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &cookie)
	var value, err = util.ParseJwt(cookie.Value)
	if err != nil {
		json.NewEncoder(writer).Encode("")
		return
	}
	id, _ := strconv.Atoi(value)
	user, err := repository.FindUserById(id)
	if err != nil {
		json.NewEncoder(writer).Encode("sai")
		return
	}
	json.NewEncoder(writer).Encode(user)
}

// func FindUserById(writer http.ResponseWriter, request *http.Request) {

// 	json.NewEncoder(writer).Encode("log out successfully")
// }
