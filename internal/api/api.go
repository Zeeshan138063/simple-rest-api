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

// New is the API constructor function
func New(db *gorm.DB) (*API, error) {
	newAPI := API{srv: new(http.Server), db: db, validate: validator.New()}
	
	// create new table upon start (will not create if one already exists)
	newAPI.db.AutoMigrate(&User{})
	
	return &newAPI, nil
}

// Run is the function for running the API server daemon
func (a *API) Run(addr string) error {
	// saving the address to server
	a.srv.Addr = addr
	
	log.Printf("running API server at http://0.0.0.0%s\n", addr)
	return a.srv.ListenAndServe()
}
