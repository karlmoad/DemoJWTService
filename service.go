package DemoJWTService

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
	router.HandleFunc("/inbox", authentication.AuthenticationHandler(msgbox.GetInboxMessages)).Methods("GET")
	router.HandleFunc("/inbox/{id}", authentication.AuthenticationHandler(msgbox.GetInboxMessage)).Methods("GET")
	router.HandleFunc("/send", authentication.AuthenticationHandler(msgbox.SendMessage)).Methods("POST")
	http.ListenAndServe(":9880", handlers.LoggingHandler(os.Stdout, router))
}



