package handler

import (
	"log"
	"net/http"

	"backend/cmd/models"

	"github.com/gorilla/mux"
)

type handler struct {
	Clients   []models.Client
	Products  []models.Product
	Images    []models.Image
	Suppliers []models.Supplier
	Addresses []models.Address
}

func NewHandler() *handler {
	return &handler{}
}

func HandleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Use(errorHandler)
	h := NewHandler()
	//  Route Handlers for Client

	myRouter.HandleFunc("/client/{client_name}", h.getClient).Methods("GET")
	myRouter.HandleFunc("/clients", h.getAllClients).Methods("GET")
	myRouter.HandleFunc("/client", h.addClient).Methods("POST")
	myRouter.HandleFunc("/client/{id}", h.updateClientAddress).Methods("PUT")
	myRouter.HandleFunc("/client/{id}", h.deleteClient).Methods("DELETE")

	//  Route Handlers for Product
	myRouter.HandleFunc("/product/{id}", h.getProduct).Methods("GET")
	myRouter.HandleFunc("/products", h.getAllProducts).Methods("GET")
	myRouter.HandleFunc("/product", h.addProduct).Methods("POST")
	myRouter.HandleFunc("/product/{id}", h.decreaseProductsAmount).Methods("PUT")
	myRouter.HandleFunc("/product/{id}", h.deleteProduct).Methods("DELETE")

	//  Route Handlers for Supplier
	myRouter.HandleFunc("/supplier/{id}", h.getSupplier).Methods("GET")
	myRouter.HandleFunc("/suppliers", h.getAllSuppliers).Methods("GET")
	myRouter.HandleFunc("/supplier", h.addSupplier).Methods("POST")
	myRouter.HandleFunc("/supplier/{id}", h.updateSupplierAddress).Methods("PUT")
	myRouter.HandleFunc("/supplier/{id}", h.deleteSupplier).Methods("DELETE")

	//  Route Handlers for Image
	myRouter.HandleFunc("/image/{id}", h.getSupplier).Methods("GET")
	myRouter.HandleFunc("/image/{product}/{image_id}", h.getImageByProductId).Methods("GET")
	myRouter.HandleFunc("/image/{}", h.addImage).Methods("POST")
	myRouter.HandleFunc("/image/{id}", h.updateImage).Methods("PUT")
	myRouter.HandleFunc("/image/{id}", h.deleteImage).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", myRouter))
}
