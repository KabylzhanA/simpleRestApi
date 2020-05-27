package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	//DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)

	DBURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", DbUser, DbPassword,DbHost, DbPort, DbName)
	println(DBURI)
	a.DB, err = gorm.Open("mysql", DBURI)
		//"root:admin@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True")
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

	a.DB.Debug().AutoMigrate(&model.Course{})

	a.DB.Exec("DROP PROCEDURE load_market_course;")

	a.DB.Exec("CREATE PROCEDURE load_market_course(" +
		" IN p_xml text," +
		" OUT retcode numeric," +
		" OUT rettext VARCHAR(255))" +
		" BEGIN" +
		" select c.retcode,c.rettext INTO RETCODE, RETTEXT" +
		" from courses c where c.p_xml = p_xml limit 1;" +
		" END;")


	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/users", a.createUser).Methods("POST")
	a.Router.HandleFunc("/users/{id}", a.deleteUser).Methods("DELETE")
	a.Router.HandleFunc("/users/{id}", a.updateUser).Methods("POST")
	a.Router.HandleFunc("/course", a.loadCourse).Methods("POST")
}

func (a *App) RunServer() {
	log.Printf("\nServer starting on port 8082")
	log.Fatal(http.ListenAndServe(":8082", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) { // this is the home route
	json.NewEncoder(w).Encode("Hello, world!")
}
