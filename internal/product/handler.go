package product

import (
	"backend/internal/apperror"
	"backend/internal/config"
	handler2 "backend/internal/handler"
	"backend/pkg/logging"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
)

type handler struct {
	logger *logging.Logger
	c      handler2.Consequences
}

func NewHandler(logger *logging.Logger) handler2.Handler {
	return &handler{
		logger: logger,
	}
}

const (
	productURL  = "/product/:uuid"
	productsURL = "/products"
)

func (h *handler) HandleRequest(cfg *config.Config) {

	myRouter := mux.NewRouter().StrictSlash(true)
	//h := NewHandler()

	//  Route Handlers for Product
	myRouter.HandleFunc(productURL, apperror.Middleware(h.getProduct)).Methods("GET")
	myRouter.HandleFunc(productsURL, apperror.Middleware(h.getAllProducts)).Methods("GET")
	myRouter.HandleFunc(productsURL, apperror.Middleware(h.addProduct)).Methods("POST")
	myRouter.HandleFunc(productURL, apperror.Middleware(h.decreaseProductsAmount)).Methods("PATCH")
	myRouter.HandleFunc(productURL, apperror.Middleware(h.deleteProduct)).Methods("DELETE")

	h.logger.Info(http.ListenAndServe(":"+cfg.Listen.Port, myRouter))
}

// ------ Product CRUD -------
// i. add product (json)
func (h *handler) addProduct(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	var prod Product
	err := json.Unmarshal(reqBody, &prod)
	if err != nil {
		return err
	}
	h.c.Products = append(h.c.Products, prod)

	fmt.Println("Endpoint Hit: addProduct")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(prod)
	if err != nil {
		return err
	}
	return nil
}

// ii. decrease available product's stock (id, amount)
func (h *handler) decreaseProductsAmount(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, prod := range h.c.Products {
		if prod.ID == id {
			h.c.Products = append(h.c.Products[:index], h.c.Products[index+1:]...)
			var prod Product
			_ = json.NewDecoder(r.Body).Decode(&prod)
			amount, err := strconv.Atoi(vars["amount"])
			if err != nil {
				return err
			}
			prod.AvailableStock -= amount
			h.c.Products = append(h.c.Products, prod)
			fmt.Println("Endpoint Hit: decreaseProductAmount")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(prod)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

// iii. get product by ID (id)
func (h *handler) getProduct(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, prod := range h.c.Products {
		if prod.ID == id {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(prod)
			if err != nil {
				return err
			}
			fmt.Println("Endpoint Hit: getProduct")
			return nil
		}
	}
	return fmt.Errorf("product by ID=%d not found", id)
}

// iv. get all products ()
func (h *handler) getAllProducts(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: getAllProducts")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(h.c.Products)
	if err != nil {
		return err
	}
	return nil
}

// v. delete product (id)
func (h *handler) deleteProduct(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, prod := range h.c.Products {
		if prod.ID == id {
			h.c.Products = append(h.c.Products[:index], h.c.Products[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			fmt.Println("Endpoint Hit: deleteProduct")
			break
		}
	}
	err := json.NewEncoder(w).Encode(h.c.Products)
	if err != nil {
		return err
	}
	return nil
}

//      ------ Product CRUD -------
