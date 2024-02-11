package handler

import (
	"backend/internal/apperror"
	"backend/internal/config"
	"backend/internal/image"
	"backend/internal/product"
	"backend/internal/supplier"
	"backend/pkg/logging"
	"log"
	"net/http"

	"backend/internal/address"
	"backend/internal/client"

	"github.com/gorilla/mux"
)

const (
	clientURLName   = "/client/{client_name}"
	clientURL       = "/client/:uuid"
	clientsURL      = "/clients"
	productURL      = "/product/:uuid"
	productsURL     = "/products"
	supplierURL     = "/supplier/:uuid"
	suppliersURL    = "/suppliers"
	imageURL        = "/image/:uuid"
	imagesURL       = "/image"
	imageProductURL = "/image/{product}/{image_id}"
)

type Handler struct {
	Clients   []client.Client
	Products  []product.Product
	Images    []image.Image
	Suppliers []supplier.Supplier
	Addresses []address.Address
}

func NewHandler() *Handler {
	return &Handler{}
}

func HandleRequests(logger *logging.Logger, cfg *config.Config) {

	myRouter := mux.NewRouter().StrictSlash(true)
	h := NewHandler()

	//  Route Handlers for Client
	myRouter.Handle(clientURLName, apperror.Middleware(h.getClient)).Methods("GET")
	myRouter.HandleFunc(clientsURL, apperror.Middleware(h.getAllClients)).Methods("GET")
	myRouter.HandleFunc(clientsURL, apperror.Middleware(h.addClient)).Methods("POST")
	myRouter.HandleFunc(clientURL, apperror.Middleware(h.updateClientAddress)).Methods("PATCH")
	myRouter.HandleFunc(clientURL, apperror.Middleware(h.deleteClient)).Methods("DELETE")

	//  Route Handlers for Product
	myRouter.HandleFunc(productURL, apperror.Middleware(h.getProduct)).Methods("GET")
	myRouter.HandleFunc(productsURL, apperror.Middleware(h.getAllProducts)).Methods("GET")
	myRouter.HandleFunc(productsURL, apperror.Middleware(h.addProduct)).Methods("POST")
	myRouter.HandleFunc(productURL, apperror.Middleware(h.decreaseProductsAmount)).Methods("PATCH")
	myRouter.HandleFunc(productURL, apperror.Middleware(h.deleteProduct)).Methods("DELETE")

	//  Route Handlers for Supplier
	myRouter.HandleFunc(supplierURL, apperror.Middleware(h.getSupplier)).Methods("GET")
	myRouter.HandleFunc(suppliersURL, apperror.Middleware(h.getAllSuppliers)).Methods("GET")
	myRouter.HandleFunc(suppliersURL, apperror.Middleware(h.addSupplier)).Methods("POST")
	myRouter.HandleFunc(supplierURL, apperror.Middleware(h.updateSupplierAddress)).Methods("PATCH")
	myRouter.HandleFunc(supplierURL, apperror.Middleware(h.deleteSupplier)).Methods("DELETE")

	//  Route Handlers for Image
	myRouter.HandleFunc(imageURL, apperror.Middleware(h.getSupplier)).Methods("GET")
	myRouter.HandleFunc(imageProductURL, apperror.Middleware(h.getImageByProductId)).Methods("GET")
	myRouter.HandleFunc(imagesURL, apperror.Middleware(h.addImage)).Methods("POST")
	myRouter.HandleFunc(imageURL, apperror.Middleware(h.updateImage)).Methods("PUT")
	myRouter.HandleFunc(imageURL, apperror.Middleware(h.deleteImage)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+cfg.Listen.Port, myRouter))
}
