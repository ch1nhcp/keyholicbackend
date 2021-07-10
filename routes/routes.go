package routes

import (
	"finalbackend/controller"
	"finalbackend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {

	//product
	r.HandleFunc("/", controller.GetCookie).Methods(http.MethodGet)
	r.HandleFunc("/", controller.SetCookie).Methods(http.MethodPost)
	r.HandleFunc("/", controller.DelCookie).Methods(http.MethodDelete)
	r.HandleFunc("/product", controller.GetAllProductPaginate).Methods(http.MethodGet)
	r.HandleFunc("/product/add", controller.AddNewProduct).Methods(http.MethodPost)
	r.HandleFunc("/product/{name}", controller.GetProductByName).Methods(http.MethodGet)
	r.HandleFunc("/category/{name}", controller.GetProductByCategory).Methods(http.MethodGet)
	// r.HandleFunc("/brand/{name}", controller.GetProductByBrand).Methods(http.MethodGet)
	//brand

	r.HandleFunc("/api/brand", middlewares.JwtVerify(controller.GetAllBrand)).Methods(http.MethodGet)
	r.HandleFunc("/api/brand/add", controller.AddNewBrand).Methods(http.MethodPost)
	r.HandleFunc("/api/brand/{id}", controller.GetBrandById).Methods(http.MethodGet)
	r.HandleFunc("/api/brand/{id}", controller.DelBrandById).Methods(http.MethodDelete)
	r.HandleFunc("/api/brand", controller.UpdateBrand).Methods(http.MethodPut)
	//user
	r.HandleFunc("/api/register", controller.Register).Methods(http.MethodGet)
	r.HandleFunc("/api/login", controller.Login).Methods(http.MethodPost)
	r.HandleFunc("/api/logout", controller.Logout).Methods(http.MethodPost)
	// auth.HandleFunc("/user/register", controller.Register).Methods(http.MethodGet)

}
