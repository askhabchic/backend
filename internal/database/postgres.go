package database

import (
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/cmd/models"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres dbname=backend port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.Address{})
	database.AutoMigrate(&models.Image{})
	database.AutoMigrate(&models.Supplier{})
	database.AutoMigrate(&models.Client{})
	database.AutoMigrate(&models.Product{})

	fmt.Println("Connect database")

	DB = database
}

func updateAddress(id string) {

}

//      ------ Client CRUD -------

// i. add client (json)
func addClient(w http.ResponseWriter, r *http.Request) {

}

// ii. delete client (id)
func deleteClient(w http.ResponseWriter, r *http.Request) {

}

// iii. get client by name and surname (name, surname)
func getClient(w http.ResponseWriter, r *http.Request) {

}

// iv. get all clients (optional: limit, offset)
func getAllClients(w http.ResponseWriter, r *http.Request) {

}

// v. update client's address (id, json - address)
func updateClientAddress(w http.ResponseWriter, r *http.Request) {

}

//      ------ Client CRUD -------
