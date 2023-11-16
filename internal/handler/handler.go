package handler

import (
	"backend/cmd/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (h *handler) updateAddress(id uuid.UUID, addr models.Address) {
	for _, address := range h.Addresses {
		if address.ID == id {
			address.City = addr.City
			address.Country = addr.Country
			address.Street = addr.Street
			break
		}
	}
}

//      ------ Client CRUD -------

// i. add client (json)
func (h *handler) addClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var client models.Client
	json.Unmarshal(reqBody, &client)

	h.Clients = append(h.Clients, client)

	fmt.Println("Endpoint Hit: addClient")
	json.NewEncoder(w).Encode(client)
}

// ii. delete client (id)
func (h *handler) deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, client := range h.Clients {
		if client.ID == uuid.Must(uuid.Parse(id)) {
			h.Clients = append(h.Clients[:index], h.Clients[index+1:]...)
			fmt.Println("Endpoint Hit: deleteClient")
			break
		}
	}
	json.NewEncoder(w).Encode(h.Clients)
}

// iii. get client by name and surname (name, surname)
func (h *handler) getClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	surname := vars["surname"]

	for _, client := range h.Clients {
		if client.Name == name && client.Surname == surname {
			fmt.Println("Endpoint Hit: getClient")
			json.NewEncoder(w).Encode(client)
			return
		}
	}
}

// iv. get all clients (optional: limit, offset)
func (h *handler) getAllClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	//limit := vars["limit"]
	//offset := vars["offset"]

	fmt.Println("Endpoint Hit: getAllClients")
	json.NewEncoder(w).Encode(h.Clients)
}

// v. update client's address (id, json - address)
func (h *handler) updateClientAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, client := range h.Clients {
		if client.ID == id {
			h.Clients = append(h.Clients[:index], h.Clients[index+1:]...)
			var addr models.Address
			_ = json.NewDecoder(r.Body).Decode(&addr)
			h.updateAddress(id, addr)
			h.Addresses = append(h.Addresses, addr)
			fmt.Println("Endpoint Hit: updateClientAddress")
			json.NewEncoder(w).Encode(addr)
			return
		}
	}
}

//      ------ Client CRUD -------

// ------ Product CRUD -------
// i. add product (json)
func (h *handler) addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product models.Product
	json.Unmarshal(reqBody, &product)
	h.Products = append(h.Products, product)

	fmt.Println("Endpoint Hit: addProduct")
	json.NewEncoder(w).Encode(product)
}

// ii. decrease available product's stock (id, amount)
func (h *handler) decreaseProductsAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, product := range h.Products {
		if product.ID == id {
			h.Products = append(h.Products[:index], h.Products[index+1:]...)
			var prod models.Product
			_ = json.NewDecoder(r.Body).Decode(&prod)
			amount, err := strconv.Atoi(vars["amount"])
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			prod.Available_stock -= amount
			h.Products = append(h.Products, prod)
			fmt.Println("Endpoint Hit: decreaseProductAmount")
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

// iii. get product by ID (id)
func (h *handler) getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, product := range h.Products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			fmt.Println("Endpoint Hit: getProduct")
			return
		}
	}
}

// iv. get all products ()
func (h *handler) getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: getAllProducts")
	json.NewEncoder(w).Encode(h.Products)
}

// v. delete product (id)
func (h *handler) deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, product := range h.Products {
		if product.ID == id {
			h.Products = append(h.Products[:index], h.Products[index+1:]...)
			fmt.Println("Endpoint Hit: deleteProduct")
			break
		}
	}
	json.NewEncoder(w).Encode(h.Products)
}

//      ------ Product CRUD -------

// ------ Supplier CRUD -------
// i. add Supplier (json)
func (h *handler) addSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var supplier models.Supplier
	json.Unmarshal(reqBody, &supplier)
	h.Suppliers = append(h.Suppliers, supplier)

	fmt.Println("Endpoint Hit: addSupplier")
	json.NewEncoder(w).Encode(supplier)
}

// ii. update Supplier's address (id, json)
func (h *handler) updateSupplierAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, supplier := range h.Suppliers {
		if supplier.ID == id {
			h.Suppliers = append(h.Suppliers[:index], h.Suppliers[index+1:]...)
			var addr models.Address
			_ = json.NewDecoder(r.Body).Decode(&addr)
			h.updateAddress(id, addr)
			h.Addresses = append(h.Addresses, addr)
			fmt.Println("Endpoint Hit: updateClientAddress")
			json.NewEncoder(w).Encode(addr)
			return
		}
	}
}

// iii. delete Supplier by ID (id)
func (h *handler) deleteSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, supplier := range h.Suppliers {
		if supplier.ID == id {
			h.Suppliers = append(h.Suppliers[:index], h.Suppliers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(h.Suppliers)
}

// iv. get all suppliers
func (h *handler) getAllSuppliers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Suppliers)
}

// v. get supplier by ID (id)
func (h *handler) getSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, supplier := range h.Suppliers {
		if supplier.ID == id {
			json.NewEncoder(w).Encode(supplier)
			return
		}
	}
}

//      ------ Supplier CRUD -------

// ------ Image CRUD -------
// i. add Image (byte array, product's id)
func (h *handler) addImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var image models.Image
	json.Unmarshal(reqBody, &image)
	h.Images = append(h.Images, image)

	json.NewEncoder(w).Encode(image)
}

// ii. update Image (image's id, string)
func (h *handler) updateImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, image := range h.Images {
		if image.ID == id {
			h.Images = append(h.Images[:index], h.Images[index+1:]...)
			var img models.Image
			_ = json.NewDecoder(r.Body).Decode(&img)
			img.Image = vars["image"]
			h.Images = append(h.Images, img)
			json.NewEncoder(w).Encode(img)
			return
		}
	}
}

// iii. delete image by ID (id)
func (h *handler) deleteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, image := range h.Images {
		if image.ID == id {
			h.Images = append(h.Images[:index], h.Images[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(h.Images)
}

// iv. get image by product's ID (id)
func (h *handler) getImageByProductId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, product := range h.Products {
		if product.ID == id {
			for _, image := range h.Images {
				if image.ID == product.Image_id {
					json.NewEncoder(w).Encode(image)
					return
				}
			}
		}
	}
}

// v. get image by image's ID
func (h *handler) getImageById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, image := range h.Images {
		if image.ID == id {
			json.NewEncoder(w).Encode(image)
			return
		}
	}
}
