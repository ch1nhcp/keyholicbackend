package routes

import (
	"finalbackend/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {
	//test cookie -ok
	r.HandleFunc("/test", controller.Test).Methods(http.MethodPost)
	r.HandleFunc("/checkcookie", controller.CheckCookie).Methods(http.MethodPost)
	//user
	// r.HandleFunc("/", controller.User).Methods(http.MethodGet)
	// r.HandleFunc("/", controller.FindUserById).Methods(http.MethodGet)
	//product -ok
	r.HandleFunc("/producthot", controller.GetProductHot).Methods(http.MethodGet)
	r.HandleFunc("/productlatest", controller.GetProductLatest).Methods(http.MethodGet)
	r.HandleFunc("/product/{name}", controller.GetProductByName).Methods(http.MethodGet)
	r.HandleFunc("/category/{name}", controller.GetProductByCategory).Methods(http.MethodGet)
	r.HandleFunc("/brand/{name}", controller.GetProductByManyBrand).Methods(http.MethodGet)

	r.HandleFunc("/product", controller.GetAllProductPaginate).Methods(http.MethodGet)
	r.HandleFunc("/product", controller.AddNewProduct).Methods(http.MethodPost)

	// brand  -ok
	r.HandleFunc("/api/brand", controller.GetAllBrand).Methods(http.MethodGet)
	r.HandleFunc("/api/brand", controller.AddNewBrand).Methods(http.MethodPost)
	r.HandleFunc("/api/brand/{id}", controller.GetBrandById).Methods(http.MethodGet)
	r.HandleFunc("/api/brand/{id}", controller.DelBrandById).Methods(http.MethodDelete)
	r.HandleFunc("/api/brand", controller.UpdateBrand).Methods(http.MethodPut)

	//user  -ok
	r.HandleFunc("/api/register", controller.Register).Methods(http.MethodPost)
	r.HandleFunc("/api/login", controller.Login).Methods(http.MethodPost)
	r.HandleFunc("/api/logout", controller.Logout).Methods(http.MethodPost)
	// r.HandleFunc("/api/cookie", controller.User).Methods(http.MethodGet)

	// category  -ok
	r.HandleFunc("/api/category", controller.GetAllCategory).Methods(http.MethodGet)
	r.HandleFunc("/api/category/add", controller.AddNewCategory).Methods(http.MethodPost)
	r.HandleFunc("/api/category/{id}", controller.GetCategoryById).Methods(http.MethodGet)
	r.HandleFunc("/api/category/{id}", controller.DelCategoryById).Methods(http.MethodDelete)
	r.HandleFunc("/api/category", controller.UpdateCategory).Methods(http.MethodPut)

	// detailproduct -ok
	r.HandleFunc("/api/detailproduct", controller.GetAllDetailproduct).Methods(http.MethodGet)
	r.HandleFunc("/api/detailproduct/add", controller.AddNewDetailproduct).Methods(http.MethodPost)
	r.HandleFunc("/api/detailproduct/{id}", controller.GetDetailproductById).Methods(http.MethodGet)
	r.HandleFunc("/api/detailproduct/{id}", controller.DelDetailproductById).Methods(http.MethodDelete)
	r.HandleFunc("/api/detailproduct", controller.UpdateDetailproduct).Methods(http.MethodPut)

	// imageproduct -ok
	r.HandleFunc("/api/imageproduct", controller.GetAllImageproduct).Methods(http.MethodGet)
	r.HandleFunc("/api/imageproduct/add", controller.AddNewImageproduct).Methods(http.MethodPost)
	r.HandleFunc("/api/imageproduct/{id}", controller.GetImageproductById).Methods(http.MethodGet)
	r.HandleFunc("/api/imageproduct/{id}", controller.DelImageproductById).Methods(http.MethodDelete)
	r.HandleFunc("/api/imageproduct", controller.UpdateImageproduct).Methods(http.MethodPut)

	// order  -ok
	r.HandleFunc("/api/order", controller.GetAllOrder).Methods(http.MethodGet)
	r.HandleFunc("/api/order", controller.AddNewOrder).Methods(http.MethodPost)
	r.HandleFunc("/api/order/{id}", controller.GetOrderById).Methods(http.MethodGet)
	r.HandleFunc("/api/order/{id}", controller.DelOrderById).Methods(http.MethodDelete)
	r.HandleFunc("/api/order", controller.UpdateOrder).Methods(http.MethodPut)

	// orderitem  -ok
	r.HandleFunc("/api/orderitem", controller.GetAllOrderitem).Methods(http.MethodGet)
	r.HandleFunc("/api/orderitem", controller.AddNewOrderitem).Methods(http.MethodPost)
	r.HandleFunc("/api/orderitem/{id}", controller.GetOrderitemById).Methods(http.MethodGet)
	r.HandleFunc("/api/orderitem/{id}", controller.DelOrderitemById).Methods(http.MethodDelete)
	r.HandleFunc("/api/orderitem", controller.UpdateOrderitem).Methods(http.MethodPut)

}
