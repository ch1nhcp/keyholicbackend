package controller

import (
	"encoding/json"
	"finalbackend/models"
	"finalbackend/repository"
	"finalbackend/util"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Register(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var User models.User
	requestBody, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(requestBody, &User)
	err := repository.RegisterUser(&User)
	if err != nil {
		json.NewEncoder(writer).Encode("user exist")
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
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	//bước 2 set cookie
	writer.Header().Set("jwt", token)
	http.SetCookie(writer, cookie)
	// writer.WriteHeader(http.StatusCreated)
	// json.NewEncoder(writer).Encode("đăng nhập thành công")
	json.NewEncoder(writer).Encode(user)
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(writer, &cookie)
	json.NewEncoder(writer).Encode("log out successfully")
}

func User(writer http.ResponseWriter, request *http.Request) {
	cookie, _ := request.Cookie("jwt")
	var value, err = util.ParseJwt(cookie.Value)
	if err != nil {
		return
	}

	json.NewEncoder(writer).Encode(value)
}
