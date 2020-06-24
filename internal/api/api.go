package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

// API holds data for running API server
type API struct {
	db       *gorm.DB
	srv      *http.Server
	validate *validator.Validate
}

// API constructor function
func New(db *gorm.DB) (*API, error) {
	newApi := API{srv: new(http.Server), db: db, validate: validator.New()}
	
	// create new table upon start (will not create if one already exists)
	newApi.db.AutoMigrate(&User{})
	
	return &newApi, nil
}

// function for running API
func (a *API) Run(addr string) error {
	// saving the address to server
	a.srv.Addr = addr
	
	log.Printf("running API server at http://0.0.0.0%s\n", addr)
	return a.srv.ListenAndServe()
}
