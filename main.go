package main

import (
	"fmt"
	"log"
	"migrations/models"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
)

type Server struct {
	DB *gorm.DB
}

func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName, DbDebug string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(DbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbDriver)
		log.Fatal("Error connecting to database: ", err)
	} else {
		fmt.Printf("Connected to %s database\n", DbDriver)
	}

	if DbDebug == "on" {
		server.DB.LogMode(true)
	}
}

func main() {
	var server = Server{}
	server.Initialize("postgres", "postgres", "root", "5432", "db", "teste", "on")

	user := models.Users{
		UserName: "Lucas " + time.Now().String(),
		Password: "123",
		Email:    "teste2@" + time.Now().String(),
	}

	createdUser, err := user.Create(server.DB) // use ID auto-increment from postgres
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	log.Println(createdUser)

}
