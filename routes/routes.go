package routes

import (
	"finalbackend/controller"
	"finalbackend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	auth := r.PathPrefix("/abc").Subrouter()
	//brand
	r.HandleFunc("/brand", controller.GetAllBrand).Methods(http.MethodGet)
	r.HandleFunc("/brand/add", controller.AddNewBrand).Methods(http.MethodPost)
	r.HandleFunc("/brand/{id}", controller.GetBrandById).Methods(http.MethodGet)
	r.HandleFunc("/brand/{id}", controller.DelBrandById).Methods(http.MethodDelete)
	r.HandleFunc("/brand", controller.UpdateBrand).Methods(http.MethodPut)
	//user
	r.HandleFunc("/user/register", controller.Register).Methods(http.MethodGet)
	r.HandleFunc("/user/login", controller.Login).Methods(http.MethodPost)
	auth.Use(middlewares.JwtVerify)

	// r.HandleFunc("/user/logout", controller.GetBrandById).Methods(http.MethodGet)
	// r.HandleFunc("/user/brand/{id}", controller.DelBrandById).Methods(http.MethodDelete)
	// r.HandleFunc("/user/brand", controller.UpdateBrand).Methods(http.MethodPut)
}
