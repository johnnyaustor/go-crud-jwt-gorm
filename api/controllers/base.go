package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	// Dialect Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/johnnyaustor/go-crud-jwt/api/models"
)

// Server Struct
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize Database
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbHost, DbPort, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		log.Println("try to connect DB Postgres")
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error: ", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	if Dbdriver == "postgres" {
		log.Println("try to connect DB Postgres")
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error: ", err)
		} else {
			log.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{})

	server.Router = mux.NewRouter()

	server.initializeRoutes()

}

// Run Application
func (server *Server) Run(addr string) {
	log.Printf("Listening to port %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
