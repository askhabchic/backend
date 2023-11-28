package database

import (
	"backend/cmd/models"
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

func ConnectDatabase(ctx context.Context) {
	dsn := "host=localhost user=postgres dbname=backend port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to postgresql")
		return
	}

	database.AutoMigrate(&models.Address{})
	database.AutoMigrate(&models.Image{})
	database.AutoMigrate(&models.Supplier{})
	database.AutoMigrate(&models.Client{})
	database.AutoMigrate(&models.Product{})

	fmt.Println("Connect database")

	DB = database
}

type repository struct {
	context.Context
	logger *logging.logger
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
