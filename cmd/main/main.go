package main

import (
	"backend/internal/config"
	"backend/internal/database"
	hh "backend/internal/handler"
	"backend/pkg/logging"
)

func main() {

	// s := Server{}
	// s.Run(":8000")
	logger := logging.GetLogger()
	logger.Info("create router in NewHandler")
	//h := hh.NewHandler()

	logger.Info("create config")
	cfg := config.GetConfig()

	logger.Info("connect to database")
	database.ConnectDatabase(cfg)

	//currentTime := time.Now()
	//h.Clients = []client.Client{
	//	client.Client{ID: uuid.New(),
	//		Name:             "Joe",
	//		Surname:          "Jonas",
	//		Birthday:         "1991-01-23",
	//		Gender:           false,
	//		RegistrationDate: currentTime.String(),
	//		AddressId:        uuid.New(),
	//	},
	//	client.Client{ID: uuid.New(),
	//		Name:             "Nick",
	//		Surname:          "Jonas",
	//		Birthday:         "1993-07-12",
	//		Gender:           false,
	//		RegistrationDate: currentTime.String(),
	//		AddressId:        uuid.New(),
	//	},
	//	client.Client{ID: uuid.New(),
	//		Name:             "Kevin",
	//		Surname:          "Jonas",
	//		Birthday:         "1990-02-04",
	//		Gender:           false,
	//		RegistrationDate: currentTime.String(),
	//		AddressId:        uuid.New(),
	//	},
	//}

	hh.HandleRequests(logger, cfg)

}
