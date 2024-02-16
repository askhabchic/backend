package client

import (
	"backend/internal/address"
	"backend/internal/apperror"
	"backend/internal/config"
	handler2 "backend/internal/handler"
	"backend/pkg/client/postgresql"
	"backend/pkg/logging"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type handler struct {
	logger    *logging.Logger
	Clients   []Client
	Addresses []address.Address
	service   Service
}

func NewHandler(logger *logging.Logger) handler2.Handler {
	return &handler{
		logger: logger,
	}
}

const (
	clientURLName = "/client/{client_name}"
	clientURL     = "/client/:uuid"
	clientsURL    = "/clients"
)

func (h *handler) HandleRequest(db *postgresql.Client, cfg *config.Config) {

	myRouter := mux.NewRouter().StrictSlash(true)

	var repository = NewClient(db, h.logger)
	NewClientService(repository)
	//  Route Handlers for Client
	myRouter.HandleFunc(clientURLName, apperror.Middleware(h.getClient)).Methods("GET")
	myRouter.HandleFunc(clientsURL, apperror.Middleware(h.getAllClients)).Methods("GET")
	myRouter.HandleFunc(clientsURL, apperror.Middleware(h.addClient)).Methods("POST")
	myRouter.HandleFunc(clientURL, apperror.Middleware(h.updateClientAddress)).Methods("PATCH")
	myRouter.HandleFunc(clientURL, apperror.Middleware(h.deleteClient)).Methods("DELETE")

	h.logger.Info(http.ListenAndServe(":"+cfg.Listen.Port, myRouter))
}

//      ------ Client CRUD -------

// i. add client (json)
func (h *handler) addClient(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	var cl Client
	err := json.Unmarshal(reqBody, &cl)
	if err != nil {
		return err
	}

	h.Clients = append(h.Clients, cl)

	h.logger.Infof("Endpoint Hit: addClient")

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(cl)
	if err != nil {
		return err
	}
	return nil
}

// ii. delete client (id)
func (h *handler) deleteClient(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, cl := range h.Clients {
		if cl.ID == uuid.Must(uuid.Parse(id)) {
			h.Clients = append(h.Clients[:index], h.Clients[index+1:]...)
			fmt.Println("Endpoint Hit: deleteClient")
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(w).Encode(h.Clients)
	if err != nil {
		return err
	}
	return nil
}

// iii. get client by name and surname (name, surname)
func (h *handler) getClient(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	surname := vars["surname"]

	for _, cli := range h.Clients {
		if cli.Name == name && cli.Surname == surname {
			fmt.Println("Endpoint Hit: getClient")
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(cli)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("client '%s %s' not found", name, surname)
}

// iv. get all clients (optional: limit, offset)
func (h *handler) getAllClients(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	//limit := vars["limit"]
	//offset := vars["offset"]

	fmt.Println("Endpoint Hit: getAllClients")

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(h.Clients)
	if err != nil {
		return err
	}
	return nil
}

// v. update client's address (id, json - address)
func (h *handler) updateClientAddress(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, cl := range h.Clients {
		if cl.ID == id {
			h.Clients = append(h.Clients[:index], h.Clients[index+1:]...)
			var addr address.Address
			err := json.NewDecoder(r.Body).Decode(&addr)
			if err != nil {
				return err
			}
			//h.updateAddress(id, addr)
			h.Addresses = append(h.Addresses, addr)
			fmt.Println("Endpoint Hit: updateClientAddress")
			w.WriteHeader(http.StatusOK)
			err = json.NewEncoder(w).Encode(addr)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

//      ------ Client CRUD -------
