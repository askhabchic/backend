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

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s",
		cfg.Storage.Host, cfg.Storage.Username, cfg.Storage.Database, cfg.Storage.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to postgresql")
		return &gorm.DB{}, err
	}

	errorMig := db.AutoMigrate(&address.Address{})
	errorMig = db.AutoMigrate(&image.Image{})
	errorMig = db.AutoMigrate(&supplier.Supplier{})
	errorMig = db.AutoMigrate(&client.Client{})
	errorMig = db.AutoMigrate(&product.Product{})
	if errorMig != nil {
		log.Fatal("Failed to AutoMigrate to postgresql")
		return &gorm.DB{}, err
	}

	fmt.Println("Connected to database. AutoMigrate is done")
	return db, nil
}
