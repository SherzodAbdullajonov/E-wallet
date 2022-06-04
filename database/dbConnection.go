package database


import (
	"fmt"
	"log"
	"os"
	
	"github.com/jinzhu/gorm"
	)
	
	var Db *gorm.DB
	var err error
	
	func Connection() {
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, dbPort)
	
	Db, err = gorm.Open(dialect, dbURI)
	if err != nil {
	log.Fatal(err)
	} else {
	fmt.Println("Successfully connected to database")
	}
	defer Db.Close()
	
	}
 	