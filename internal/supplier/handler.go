package supplier

import (
	"backend/internal/address"
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
	supplierURL  = "/supplier/:uuid"
	suppliersURL = "/suppliers"
)

func (h *handler) HandleRequest(cfg *config.Config) {
	myRouter := mux.NewRouter().StrictSlash(true)
	//h := NewHandler()

	//  Route Handlers for Supplier
	myRouter.HandleFunc(supplierURL, apperror.Middleware(h.getSupplier)).Methods("GET")
	myRouter.HandleFunc(suppliersURL, apperror.Middleware(h.getAllSuppliers)).Methods("GET")
	myRouter.HandleFunc(suppliersURL, apperror.Middleware(h.addSupplier)).Methods("POST")
	myRouter.HandleFunc(supplierURL, apperror.Middleware(h.updateSupplierAddress)).Methods("PATCH")
	myRouter.HandleFunc(supplierURL, apperror.Middleware(h.deleteSupplier)).Methods("DELETE")

	h.logger.Info(http.ListenAndServe(":"+cfg.Listen.Port, myRouter))
}

// ------ Supplier CRUD -------
// i. add Supplier (json)
func (h *handler) addSupplier(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	var suppl Supplier
	err := json.Unmarshal(reqBody, &suppl)
	if err != nil {
		return err
	}
	h.c.Suppliers = append(h.c.Suppliers, suppl)

	fmt.Println("Endpoint Hit: addSupplier")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(suppl)
	if err != nil {
		return err
	}
	return nil
}

// ii. update Supplier's address (id, json)
func (h *handler) updateSupplierAddress(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, suppl := range h.c.Suppliers {
		if suppl.ID == id {
			h.c.Suppliers = append(h.c.Suppliers[:index], h.c.Suppliers[index+1:]...)
			var addr address.Address
			_ = json.NewDecoder(r.Body).Decode(&addr)
			//h.updateAddress(id, addr)
			h.c.Addresses = append(h.c.Addresses, addr)
			fmt.Println("Endpoint Hit: updateClientAddress")
			w.WriteHeader(http.StatusNoContent)
			err := json.NewEncoder(w).Encode(addr)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

// iii. delete Supplier by ID (id)
func (h *handler) deleteSupplier(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, suppl := range h.c.Suppliers {
		if suppl.ID == id {
			h.c.Suppliers = append(h.c.Suppliers[:index], h.c.Suppliers[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(w).Encode(h.c.Suppliers)
	if err != nil {
		return err
	}
	return nil
}

// iv. get all suppliers
func (h *handler) getAllSuppliers(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(h.c.Suppliers)
	if err != nil {
		return err
	}
	return nil
}

// v. get supplier by ID (id)
func (h *handler) getSupplier(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, suppl := range h.c.Suppliers {
		if suppl.ID == id {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(suppl)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("supplier by ID=%d not found", id)
}

//      ------ Supplier CRUD -------
