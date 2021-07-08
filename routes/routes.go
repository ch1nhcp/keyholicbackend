package routes

import (
	"finalbackend/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	//brand
	r.HandleFunc("/brand", controller.GetAllBrand).Methods(http.MethodGet)
	r.HandleFunc("/brand/add", controller.AddNewBrand).Methods(http.MethodPost)
	r.HandleFunc("/brand/{id}", controller.GetBrandById).Methods(http.MethodGet)
	r.HandleFunc("/brand/{id}", controller.DelBrandById).Methods(http.MethodDelete)
	r.HandleFunc("/brand", controller.UpdateBrand).Methods(http.MethodPut)
	//category
}
