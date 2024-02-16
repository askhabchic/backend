package client

import (
	"backend/internal/client"
	"backend/pkg/logging"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type Repository struct {
	psgr   *gorm.DB
	logger *logging.Logger
}

func NewRepository(client *gorm.DB, logger *logging.Logger) *Repository {
	return &Repository{
		psgr:   client,
		logger: logger,
	}
}

//TODO связать с client/handler

func (r Repository) Create(ctx context.Context, cl *client.Client) error {
	q := `INSERT INTO client (client_name, client_surname, birthday, gender, registration_date, address_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	queryRow := r.psgr.QueryRow(ctx, q, cl.Name, cl.Surname, cl.Birthday, cl.Gender, cl.RegistrationDate, cl.AddressId)
	if err := queryRow.Scan(&cl.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r Repository) FindAll(ctx context.Context, limit, offset int) (clients []client.Client, err error) {
	q := `SELECT id, client_name, client_surname, birthday, gender, registration_date, address_id FROM public.client LIMIT $1 OFFSET`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))
	rows, err := r.psgr.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	clients = make([]client.Client, 0)

	for rows.Next() {
		var cl client.Client

		err := rows.Scan(&cl.ID, &cl.Name)
		if err != nil {
			return nil, err
		}

		clients = append(clients, cl)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}

func (r Repository) FindOne(ctx context.Context, id string) (client.Client, error) {
	q := `SELECT id, client_name, client_surname, birthday, gender, registration_date, address_id FROM public.client WHERE id = $1`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	var cl client.Client
	queryRow := r.psgr.QueryRow(ctx, q, id)
	err := queryRow.Scan(&cl.ID, &cl.Name)
	if err != nil {
		return client.Client{}, err
	}
	return cl, nil
}

func (r Repository) Update(ctx context.Context, cl client.Client, args ...interface{}) error {
	q := `UPDATE client SET address_id = $1 WHERE id = $2`

	queryRow := r.psgr.QueryRow(ctx, q, args, cl.ID)
	if err := queryRow.Scan(&cl.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM client WHERE id = $1`

	queryRow := r.psgr.QueryRow(ctx, q, id)
	err := queryRow.Scan(id)
	if err != nil {
		return err
	}
	return nil
}
