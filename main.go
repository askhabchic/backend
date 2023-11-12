package main

import (
	"backend/internal/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Article ...
type Article struct {
	ID     string `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"author"`
	Link   string `json:"link"`
}

type Client struct {
	ID                string `json:"id"`
	Name              string `json:"client_name"`
	Surname           string `json:"client_surname"`
	Birthday          string `json:"birthday"`
	Gender            string `json:"gender"`
	Registration_date string `json:"registration_date"`
	Address_id        string `json:"address_id"`
}

type Product struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Category         string `json:"category"`
	Price            string `json:"price"`
	Available_stock  int    `json:"available_stock"`
	Last_update_date string `json:"last_update_date"`
	Supplier_id      string `json:"supplier_id"`
	Image_id         string `json:"image_id"`
}

type Image struct {
	ID    string `json:"id"`
	Image string `json:"image"`
}

type Supplier struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Address_id   string `json:"address_id"`
	Phone_number string `json:"phone_number"`
}

type Address struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
}

// Articles ...
var Articles []Article

var Clients []Client
var Products []Product
var Images []Image
var Suppliers []Supplier
var Addresses []Address

func updateAddress(id string) {
	for index, address := range Addresses {
		if address.ID == id {
			Addresses = append(Addresses[:index], Addresses[index+1:]...)
			break
		}
	}
}

//      ------ Client CRUD -------

// i. add client (json)
func addClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var client Client
	json.Unmarshal(reqBody, &client)
	Clients = append(Clients, client)

	fmt.Println("Endpoint Hit: addClient")
	json.NewEncoder(w).Encode(client)
}

// ii. delete client (id)
func deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	for index, client := range Clients {
		if client.ID == id {
			Clients = append(Clients[:index], Clients[index+1:]...)
			fmt.Println("Endpoint Hit: deleteClient")
			break
		}
	}
	json.NewEncoder(w).Encode(Clients)
}

// iii. get client by name and surname (name, surname)
func getClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	name := vars["name"]
	surname := vars["surname"]

	for _, client := range Clients {
		if client.Name == name && client.Surname == surname {
			fmt.Println("Endpoint Hit: getClient")
			json.NewEncoder(w).Encode(client)
			return
		}
	}
}

// iv. get all clients (optional: limit, offset)
func getAllClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//vars := mux.Vars(r)
	//limit := vars["limit"]
	//offset := vars["offset"]

	fmt.Println("Endpoint Hit: getAllClients")
	json.NewEncoder(w).Encode(Clients)
}

// v. update client's address (id, json - address)
func updateClientAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for index, client := range Clients {
		if client.ID == vars["id"] {
			Clients = append(Clients[:index], Clients[index+1:]...)
			var cl Client
			_ = json.NewDecoder(r.Body).Decode(&cl)
			cl.Address_id = vars["address"]
			updateAddress(vars["id"])
			Clients = append(Clients, cl)
			fmt.Println("Endpoint Hit: updateClientAddress")
			json.NewEncoder(w).Encode(cl)
			return
		}
	}
}

//      ------ Client CRUD -------

// ------ Product CRUD -------
// i. add product (json)
func addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var product Product
	json.Unmarshal(reqBody, &product)
	Products = append(Products, product)

	fmt.Println("Endpoint Hit: addProduct")
	json.NewEncoder(w).Encode(product)
}

// ii. decrease available product's stock (id, amount)
func decreaseProductsAmount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for index, product := range Products {
		if product.ID == vars["id"] {
			Products = append(Products[:index], Products[index+1:]...)
			var prod Product
			_ = json.NewDecoder(r.Body).Decode(&prod)
			amount, err := strconv.Atoi(vars["amount"])
			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
			prod.Available_stock -= amount
			Products = append(Products, prod)
			fmt.Println("Endpoint Hit: decreaseProductAmount")
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

// iii. get product by ID (id)
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for _, product := range Products {
		if product.ID == vars["id"] {
			json.NewEncoder(w).Encode(product)
			fmt.Println("Endpoint Hit: getProduct")
			return
		}
	}
}

// iv. get all products ()
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: getAllProducts")
	json.NewEncoder(w).Encode(Products)
}

// v. delete product (id)
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for index, product := range Products {
		if product.ID == vars["id"] {
			Products = append(Products[:index], Products[index+1:]...)
			fmt.Println("Endpoint Hit: deleteProduct")
			break
		}
	}
	json.NewEncoder(w).Encode(Products)
}

//      ------ Product CRUD -------

// ------ Supplier CRUD -------
// i. add Supplier (json)
func addSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var supplier Supplier
	json.Unmarshal(reqBody, &supplier)
	Suppliers = append(Suppliers, supplier)

	fmt.Println("Endpoint Hit: addSupplier")
	json.NewEncoder(w).Encode(supplier)
}

// ii. update Supplier's address (id, json)
func updateSupplierAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for index, supplier := range Suppliers {
		if supplier.ID == vars["id"] {
			Suppliers = append(Suppliers[:index], Suppliers[index+1:]...)
			var sup Supplier
			_ = json.NewDecoder(r.Body).Decode(&sup)
			sup.Address_id = vars["address"]
			updateAddress(vars["id"])
			Suppliers = append(Suppliers, sup)
			json.NewEncoder(w).Encode(sup)
			return
		}
	}
}

// iii. delete Supplier by ID (id)
func deleteSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for index, supplier := range Suppliers {
		if supplier.ID == vars["id"] {
			Suppliers = append(Suppliers[:index], Suppliers[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Suppliers)
}

// iv. get all suppliers
func getAllSuppliers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Suppliers)
}

// v. get supplier by ID (id)
func getSupplier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for _, supplier := range Suppliers {
		if supplier.ID == vars["id"] {
			json.NewEncoder(w).Encode(supplier)
			return
		}
	}
}

//      ------ Supplier CRUD -------

// ------ Image CRUD -------
// i. add Image (byte array, product's id)
func addImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var image Image
	json.Unmarshal(reqBody, &image)
	Images = append(Images, image)

	json.NewEncoder(w).Encode(image)
}

// ii. update Image (image's id, string)
func updateImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)

	for index, image := range Images {
		if image.ID == vars["id"] {
			Images = append(Images[:index], Images[index+1:]...)
			var img Image
			_ = json.NewDecoder(r.Body).Decode(&img)
			img.Image = vars["image"]
			Images = append(Images, img)
			json.NewEncoder(w).Encode(img)
			return
		}
	}
}

// iii. delete image by ID (id)
func deleteImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)

	for index, image := range Images {
		if image.ID == vars["id"] {
			Images = append(Images[:index], Images[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Images)
}

// iv. get image by product's ID (id)
func getImageByProductId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)

	for _, product := range Products {
		if product.ID == vars["id"] {
			for _, image := range Images {
				if image.ID == product.Image_id {
					json.NewEncoder(w).Encode(image)
					return
				}
			}
		}
	}
}

// v. get image by image's ID
func getImageById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)

	for _, image := range Images {
		if image.ID == vars["id"] {
			json.NewEncoder(w).Encode(image)
			return
		}
	}
}

//      ------ Image CRUD -------

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// func handleRequests() {
//     myRouter := mux.NewRouter().StrictSlash(true)
//     myRouter.HandleFunc("/", homePage)
//     myRouter.HandleFunc("/articles", returnAllArticles)
//     myRouter.HandleFunc("/article/{id}",returnSingleArticle)
//     myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
//     myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
//     log.Fatal(http.ListenAndServe(":8000", myRouter))
// }

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	//  Route Handlers for Client
	myRouter.HandleFunc("/client/{client_name}", getClient).Methods("GET")
	myRouter.HandleFunc("/clients", getAllClients).Methods("GET")
	myRouter.HandleFunc("/client", addClient).Methods("POST")
	myRouter.HandleFunc("/client/{id}", updateClientAddress).Methods("PUT")
	myRouter.HandleFunc("/client/{id}", deleteClient).Methods("DELETE")

	//  Route Handlers for Product
	myRouter.HandleFunc("/product/{id}", getProduct).Methods("GET")
	myRouter.HandleFunc("/products", getAllProducts).Methods("GET")
	myRouter.HandleFunc("/product", addProduct).Methods("POST")
	myRouter.HandleFunc("/product/{id}", decreaseProductsAmount).Methods("PUT")
	myRouter.HandleFunc("/product/{id}", deleteProduct).Methods("DELETE")

	//  Route Handlers for Supplier
	myRouter.HandleFunc("/supplier/{id}", getSupplier).Methods("GET")
	myRouter.HandleFunc("/suppliers", getAllSuppliers).Methods("GET")
	myRouter.HandleFunc("/supplier", addSupplier).Methods("POST")
	myRouter.HandleFunc("/supplier/{id}", updateSupplierAddress).Methods("PUT")
	myRouter.HandleFunc("/supplier/{id}", deleteSupplier).Methods("DELETE")

	//  Route Handlers for Image
	myRouter.HandleFunc("/image/{id}", getSupplier).Methods("GET")
	myRouter.HandleFunc("/image/{product}/{image_id}", getImageByProductId).Methods("GET")
	myRouter.HandleFunc("/image/{}", addImage).Methods("POST")
	myRouter.HandleFunc("/image/{id}", updateImage).Methods("PUT")
	myRouter.HandleFunc("/image/{id}", deleteImage).Methods("DELETE")

	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.ID == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func main() {

	// s := Server{}
	// s.Run(":8000")
	// Articles = []Article{
	// 	Article{ID: "1",
	// 		Title: "Python Intermediate and Advanced 101",
	// 		Author: "Arkaprabha Majumdar",
	// 		Link:   "https://www.amazon.com/dp/B089KVK23P"},
	// 	Article{ID: "2",
	// 		Title: "R programming Advanced",
	// 		Author: "Arkaprabha Majumdar",
	// 		Link:   "https://www.amazon.com/dp/B089WH12CR"},
	// 	Article{ID: "3",
	// 		Title: "R programming Fundamentals",
	// 		Author: "Arkaprabha Majumdar",
	// 		Link:   "https://www.amazon.com/dp/B089S58WWG"},
	// }

	database.ConnectDatabase()

	currentTime := time.Now()
	Clients = []Client{
		Client{ID: "fa82807a-66d6-11ee-8c99-0242ac120002",
			Name:              "Joe",
			Surname:           "Jonas",
			Birthday:          "1991-01-23",
			Gender:            "male",
			Registration_date: currentTime.String(),
			Address_id:        "0239d3c2-66d7-11ee-8c99-0242ac120002",
		},
		Client{ID: "36c0ec84-66d7-11ee-8c99-0242ac120002",
			Name:              "Nick",
			Surname:           "Jonas",
			Birthday:          "1993-07-12",
			Gender:            "male",
			Registration_date: currentTime.String(),
			Address_id:        "3ed74c38-66d7-11ee-8c99-0242ac120002",
		},
		Client{ID: "664803b6-66d7-11ee-8c99-0242ac120002",
			Name:              "Kevin",
			Surname:           "Jonas",
			Birthday:          "1990-02-04",
			Gender:            "male",
			Registration_date: currentTime.String(),
			Address_id:        "6b98fdc0-66d7-11ee-8c99-0242ac120002",
		},
	}
	handleRequests()
}
