package handler

import (
<<<<<<< HEAD
	"backend/internal/config"
	"backend/pkg/client/postgresql"
)

type Handler interface {
	HandleRequest(db *postgresql.Client, cfg *config.Config)
}

//func (h *Handler) updateAddress(id uuid.UUID, addr address.Address) error {
//	for _, add := range h.Addresses {
//		if add.ID == id {
//			add.City = addr.City
//			add.Country = addr.Country
//			add.Street = addr.Street
//			break
//		}
//	}
//	return nil
//}
=======
	"backend/internal/address"
	"backend/internal/client"
	"backend/internal/image"
	"backend/internal/product"
	"backend/internal/supplier"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *Handler) updateAddress(id uuid.UUID, addr address.Address) error {
	for _, add := range h.Addresses {
		if add.ID == id {
			add.City = addr.City
			add.Country = addr.Country
			add.Street = addr.Street
			break
		}
	}
	return nil
}

//      ------ Client CRUD -------

// i. add client (json)
func (h *Handler) addClient(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	var cl client.Client
	err := json.Unmarshal(reqBody, &cl)
	if err != nil {
		return err
	}

	h.Clients = append(h.Clients, cl)

	fmt.Println("Endpoint Hit: addClient")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(cl)
	if err != nil {
		return err
	}
	return nil
}

// ii. delete client (id)
func (h *Handler) deleteClient(w http.ResponseWriter, r *http.Request) error {
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
func (h *Handler) getClient(w http.ResponseWriter, r *http.Request) error {
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
func (h *Handler) getAllClients(w http.ResponseWriter, r *http.Request) error {
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
func (h *Handler) updateClientAddress(w http.ResponseWriter, r *http.Request) error {
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
			h.updateAddress(id, addr)
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

// ------ Product CRUD -------
// i. add product (json)
func (h *Handler) addProduct(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	var prod product.Product
	err := json.Unmarshal(reqBody, &prod)
	if err != nil {
		return err
	}
	h.Products = append(h.Products, prod)

	fmt.Println("Endpoint Hit: addProduct")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(prod)
	if err != nil {
		return err
	}
	return nil
}

// ii. decrease available product's stock (id, amount)
func (h *Handler) decreaseProductsAmount(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, prod := range h.Products {
		if prod.ID == id {
			h.Products = append(h.Products[:index], h.Products[index+1:]...)
			var prod product.Product
			_ = json.NewDecoder(r.Body).Decode(&prod)
			amount, err := strconv.Atoi(vars["amount"])
			if err != nil {
				return err
			}
			prod.AvailableStock -= amount
			h.Products = append(h.Products, prod)
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
func (h *Handler) getProduct(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, prod := range h.Products {
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
func (h *Handler) getAllProducts(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: getAllProducts")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(h.Products)
	if err != nil {
		return err
	}
	return nil
}

// v. delete product (id)
func (h *Handler) deleteProduct(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, prod := range h.Products {
		if prod.ID == id {
			h.Products = append(h.Products[:index], h.Products[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			fmt.Println("Endpoint Hit: deleteProduct")
			break
		}
	}
	err := json.NewEncoder(w).Encode(h.Products)
	if err != nil {
		return err
	}
	return nil
}

//      ------ Product CRUD -------

// ------ Supplier CRUD -------
// i. add Supplier (json)
func (h *Handler) addSupplier(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := io.ReadAll(r.Body)
	var suppl supplier.Supplier
	err := json.Unmarshal(reqBody, &suppl)
	if err != nil {
		return err
	}
	h.Suppliers = append(h.Suppliers, suppl)

	fmt.Println("Endpoint Hit: addSupplier")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(suppl)
	if err != nil {
		return err
	}
	return nil
}

// ii. update Supplier's address (id, json)
func (h *Handler) updateSupplierAddress(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, suppl := range h.Suppliers {
		if suppl.ID == id {
			h.Suppliers = append(h.Suppliers[:index], h.Suppliers[index+1:]...)
			var addr address.Address
			_ = json.NewDecoder(r.Body).Decode(&addr)
			h.updateAddress(id, addr)
			h.Addresses = append(h.Addresses, addr)
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
func (h *Handler) deleteSupplier(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, suppl := range h.Suppliers {
		if suppl.ID == id {
			h.Suppliers = append(h.Suppliers[:index], h.Suppliers[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(w).Encode(h.Suppliers)
	if err != nil {
		return err
	}
	return nil
}

// iv. get all suppliers
func (h *Handler) getAllSuppliers(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(h.Suppliers)
	if err != nil {
		return err
	}
	return nil
}

// v. get supplier by ID (id)
func (h *Handler) getSupplier(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, suppl := range h.Suppliers {
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

// ------ Image CRUD -------
// i. add Image (byte array, product's id)
func (h *Handler) addImage(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	reqBody, _ := io.ReadAll(r.Body)
	var img image.Image
	err := json.Unmarshal(reqBody, &img)
	if err != nil {
		return err
	}
	h.Images = append(h.Images, img)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(img)
	if err != nil {
		return err
	}
	return nil
}

// ii. update Image (image's id, string)
func (h *Handler) updateImage(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, img := range h.Images {
		if img.ID == id {
			h.Images = append(h.Images[:index], h.Images[index+1:]...)
			var img image.Image
			err := json.NewDecoder(r.Body).Decode(&img)
			if err != nil {
				return err
			}
			img.Image = vars["image"]
			h.Images = append(h.Images, img)
			w.WriteHeader(http.StatusNoContent)
			err = json.NewEncoder(w).Encode(img)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

// iii. delete image by ID (id)
func (h *Handler) deleteImage(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, img := range h.Images {
		if img.ID == id {
			h.Images = append(h.Images[:index], h.Images[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(w).Encode(h.Images)
	if err != nil {
		return err
	}
	return nil
}

// iv. get image by product's ID (id)
func (h *Handler) getImageByProductId(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, prod := range h.Products {
		if prod.ID == id {
			for _, img := range h.Images {
				if img.ID == prod.ImageId {
					err := json.NewEncoder(w).Encode(img)
					if err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	return fmt.Errorf("image by Product ID=%d not found", id)
}

// v. get image by image's ID
func (h *Handler) getImageById(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, img := range h.Images {
		if img.ID == id {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(img)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("image by ID=%d not found", id)
}
>>>>>>> ac96bf5b98ac830e992a085f7399d2bb1bda3c6f
