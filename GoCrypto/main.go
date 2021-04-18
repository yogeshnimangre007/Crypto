package main 

import (
	"log"
	"net/http"
	"time"

	"Users/Downloads/GoCrypto/Controller"
	"Users/Downloads/GoCrypto/Model"
	"github.com/gorilla/mux"
	"github.com/labstack/echo"//Used for helping in unmarshalling of timestamps
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).
		Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).
		Methods("POST")
	r.HandleFunc("/profile", controller.ProfileHandler).
		Methods("GET")

	e := echo.New()

	e.POST("/User", CreateEvent)

	if err := e.Start("9000"); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":9000", r))
}
