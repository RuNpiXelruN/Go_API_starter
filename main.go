package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"../app/src/controller"
	"../app/src/model"
)

func main() {
	controller.Startup()
	db := connectToDatabase()
	defer db.Close()
	http.ListenAndServe(":8080", nil)
}

func connectToDatabase() *gorm.DB {
	db, err := gorm.Open("postgres", "user=gorm password=gorm dbname=gorm sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
}
