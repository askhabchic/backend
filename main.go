package main

import (
	"backend/cmd/models"
	"backend/internal/database"
	hh "backend/internal/handler"
	"time"

	"github.com/google/uuid"
)

func main() {

	// s := Server{}
	// s.Run(":8000")
	h := hh.NewHendler()
	database.ConnectDatabase()

	currentTime := time.Now()
	h.Clients = []models.Client{
		models.Client{ID: uuid.New(),
			Name:              "Joe",
			Surname:           "Jonas",
			Birthday:          "1991-01-23",
			Gender:            false,
			Registration_date: currentTime.String(),
			Address_id:        uuid.New(),
		},
		models.Client{ID: uuid.New(),
			Name:              "Nick",
			Surname:           "Jonas",
			Birthday:          "1993-07-12",
			Gender:            false,
			Registration_date: currentTime.String(),
			Address_id:        uuid.New(),
		},
		models.Client{ID: uuid.New(),
			Name:              "Kevin",
			Surname:           "Jonas",
			Birthday:          "1990-02-04",
			Gender:            false,
			Registration_date: currentTime.String(),
			Address_id:        uuid.New(),
		},
	}
	hh.HandleRequests()
}
