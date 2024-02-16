package postgresql

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
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func DoWithTries(fn func() error, attempts int, delay time.Duration) (err error) {
	for attempts > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attempts--

			continue
		}

		return nil
	}
	return
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

func NewClient(ctx context.Context, maxAttempts int, cfg *config.Config) (*gorm.DB, error) {
	var pool *gorm.DB
	var err error
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Storage.Username,
		cfg.Storage.Password, cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.Database)
	err = DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			fmt.Println("failed to connect to postgresql")
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error connect to database with timeout")
	}
	return pool, nil
}
