package gojira

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "gojira api")
}

// App - Main entry point.
func App() {
	log.Info("starting app")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/ticket", createTicket).Methods("POST")
	router.HandleFunc("/tickets", getAllTickets).Methods("GET")
	router.HandleFunc("/tickets/{id}", getOneTicket).Methods("GET")
	router.HandleFunc("/tickets/{id}", updateTicket).Methods("PATCH")
	router.HandleFunc("/tickets/{id}", deleteTicket).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
