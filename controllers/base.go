package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"simpleRestApi/model"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(DbHost, DbPort, DbUser, DbName, DbPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to database %s", DbName)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database %s", DbName)
	}

	if !a.DB.HasTable(&model.User{}) {
		a.DB.Debug().AutoMigrate(&model.User{})
		a.DB.Exec("ALTER TABLE users add column birthday date")
	} else {
		a.DB.Debug().AutoMigrate(&model.User{})
	}

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/users", a.createUser).Methods("POST")
	a.Router.HandleFunc("/users/{id}", a.deleteUser).Methods("DELETE")
	a.Router.HandleFunc("/users/{id}", a.updateUser).Methods("POST")
}

func (a *App) RunServer() {
	log.Printf("\nServer starting on port 8082")
	log.Fatal(http.ListenAndServe(":8082", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) { // this is the home route
	json.NewEncoder(w).Encode("Hello, world!")
}
