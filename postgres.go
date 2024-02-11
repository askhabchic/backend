package database

import (
	"backend/internal/address"
	"backend/internal/client"
	"backend/internal/config"
	"backend/internal/image"
	"backend/internal/product"
	"backend/internal/supplier"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Connection interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	// Begin starts a transaction block from the *Conn without explicitly setting a transaction mode (see BeginTx with TxOptions if transaction mode is required).
	Begin(ctx context.Context) (pgx.Tx, error)
}

var DB *gorm.DB

func ConnectDatabase(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s",
		cfg.Storage.Host, cfg.Storage.Username, cfg.Storage.Database, cfg.Storage.Port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to postgresql")
		return
	}

	errorMig := database.AutoMigrate(&address.Address{})
	if errorMig != nil {
		return
	}
	errorMig = database.AutoMigrate(&image.Image{})
	if errorMig != nil {
		return
	}
	errorMig = database.AutoMigrate(&supplier.Supplier{})
	if errorMig != nil {
		return
	}
	errorMig = database.AutoMigrate(&client.Client{})
	if errorMig != nil {
		return
	}
	errorMig = database.AutoMigrate(&product.Product{})
	if errorMig != nil {
		return
	}

	fmt.Println("Connect to database")

	DB = database
}

type repository struct {
	context.Context
	//logger *logging.logger
	// Pool
}

func (r *repository) updateAddress(id string) {

}

//      ------ Client CRUD -------

// i. add client (json)
func addClient() {

}

// ii. delete client (id)
func deleteClient() {

}

// iii. get client by name and surname (name, surname)
func getClient() {

}

// iv. get all clients (optional: limit, offset)
func getAllClients() {

}

// v. update client's address (id, json - address)
func updateClientAddress() {

}

//      ------ Client CRUD -------
