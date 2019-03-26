package main

import (
	"fmt"
	"github.com/MrTBit/restapi/app"
	"github.com/MrTBit/restapi/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	//router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	//router.HandleFunc("/api/user/{id}/contacts", controllers.GetContactsFor).Methods("GET") // user/2/contacts

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT") //get port from .env, we didn't speicfy so it will return empty string
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //launch the app, localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}

}
