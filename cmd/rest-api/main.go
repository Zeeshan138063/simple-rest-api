package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

func main() {
	// load environment variables for secured data
	port := os.Getenv("PORT")
	
	// default port
	if port == "" {
		port = "5400"
	}
	
	// open database connection
	db, err := gorm.Open("sqlite3", "db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	// close db connection upon application closes
	defer db.Close()
}
