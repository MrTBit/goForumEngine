package main

import (
	"fmt"
	"github.com/MrTBit/goForumEngine/app"
	"github.com/MrTBit/goForumEngine/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/boards/new", controllers.CreateBoard).Methods("POST")
	router.HandleFunc("/api/boards/{boardID}/new", controllers.CreateThread).Methods("POST")
	router.HandleFunc("/api/boards/{boardID}/threads/{threadID}/new", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/api/boards", controllers.GetBoards).Methods("GET")
	router.HandleFunc("/api/boards/{boardID}/threads", controllers.GetThreadsFor).Methods("GET")
	router.HandleFunc("/api/boards/{boardID}/threads/{threadID}", controllers.GetThread).Methods("GET")
	router.HandleFunc("/api/boards/{boardID}/threads/{threadID}/posts", controllers.GetCommentsFor).Methods("GET")
	router.HandleFunc("/api/boards/{boardID}/threads/{threadID}/posts/{postID}", controllers.GetComment).Methods("GET")

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
