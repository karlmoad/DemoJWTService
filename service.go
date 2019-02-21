package main

import (
	"DemoJWTService/authentication"
	"DemoJWTService/msgbox"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	router := mux.NewRouter()
	router.HandleFunc("/secret", authentication.AuthenticationHandler(msgbox.GetSecretMessage)).Methods("GET")
	router.HandleFunc("/secret/{count}", authentication.AuthenticationHandler(msgbox.GetSecretMessage)).Methods("GET")

	// CORS options
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization","Content-Length", "Accept"})
	originsOk := handlers.AllowedOrigins([]string{"*"})  // allow all inbound domains
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	http.ListenAndServe(":30200", handlers.CORS(headersOk, originsOk, methodsOk)(handlers.LoggingHandler(os.Stdout, router)))
}



