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
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))
	if err != nil {
		json.NewEncoder(writer).Encode(user)
		return
	}
	//đưa vào cookie
	//b1 : tọa cookie
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	//bước 2 set cookie
	http.SetCookie(writer, &cookie)
	// writer.WriteHeader(http.StatusCreated)
	// json.NewEncoder(writer).Encode("đăng nhập thành công")
	json.NewEncoder(writer).Encode(user)

}
