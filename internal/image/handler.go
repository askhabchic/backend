package image

import (
	"backend/internal/apperror"
	"backend/internal/config"
	handler2 "backend/internal/handler"
	"backend/pkg/logging"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type handler struct {
	logger *logging.Logger
	c      handler2.Consequences
}

func NewHandler(logger *logging.Logger) handler2.Handler {
	return &handler{
		logger: logger,
	}
}

const (
	imageURL        = "/image/:uuid"
	imagesURL       = "/image"
	imageProductURL = "/image/{product}/{image_id}"
)

func (h *handler) HandleRequest(cfg *config.Config) {

	myRouter := mux.NewRouter().StrictSlash(true)
	//h := NewHandler()

	//  Route Handlers for Image
	//myRouter.HandleFunc(imageURL, apperror.Middleware(h.getSupplier)).Methods("GET")
	myRouter.HandleFunc(imageProductURL, apperror.Middleware(h.getImageByProductId)).Methods("GET")
	myRouter.HandleFunc(imagesURL, apperror.Middleware(h.addImage)).Methods("POST")
	myRouter.HandleFunc(imageURL, apperror.Middleware(h.updateImage)).Methods("PUT")
	myRouter.HandleFunc(imageURL, apperror.Middleware(h.deleteImage)).Methods("DELETE")

	h.logger.Info(http.ListenAndServe(":"+cfg.Listen.Port, myRouter))
}

// ------ Image CRUD -------
// i. add Image (byte array, product's id)
func (h *handler) addImage(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	reqBody, _ := io.ReadAll(r.Body)
	var img Image
	err := json.Unmarshal(reqBody, &img)
	if err != nil {
		return err
	}
	h.c.Images = append(h.c.Images, img)
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(img)
	if err != nil {
		return err
	}
	return nil
}

// ii. update Image (image's id, string)
func (h *handler) updateImage(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, img := range h.c.Images {
		if img.ID == id {
			h.c.Images = append(h.c.Images[:index], h.c.Images[index+1:]...)
			var img Image
			err := json.NewDecoder(r.Body).Decode(&img)
			if err != nil {
				return err
			}
			img.Image = vars["image"]
			h.c.Images = append(h.c.Images, img)
			w.WriteHeader(http.StatusNoContent)
			err = json.NewEncoder(w).Encode(img)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

// iii. delete image by ID (id)
func (h *handler) deleteImage(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for index, img := range h.c.Images {
		if img.ID == id {
			h.c.Images = append(h.c.Images[:index], h.c.Images[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
	err := json.NewEncoder(w).Encode(h.c.Images)
	if err != nil {
		return err
	}
	return nil
}

// iv. get image by product's ID (id)
func (h *handler) getImageByProductId(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, prod := range h.c.Products {
		if prod.ID == id {
			for _, img := range h.c.Images {
				if img.ID == prod.ImageId {
					err := json.NewEncoder(w).Encode(img)
					if err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	return fmt.Errorf("image by Product ID=%d not found", id)
}

// v. get image by image's ID
func (h *handler) getImageById(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	vars := mux.Vars(r)
	id := uuid.Must(uuid.Parse(vars["id"]))

	for _, img := range h.c.Images {
		if img.ID == id {
			w.WriteHeader(http.StatusOK)
			err := json.NewEncoder(w).Encode(img)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("image by ID=%d not found", id)
}
