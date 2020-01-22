package gojira

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// dummy datasource for now
var (
	tickets = []Ticket{}
)

func createTicket(w http.ResponseWriter, r *http.Request) {
	var ticket Ticket
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		// TODO: meaningful response message
	}

	id := len(tickets) + 1
	json.Unmarshal(reqBody, &ticket)
	ticket.ID = id

	tickets = append(tickets, ticket)

	log.WithFields(log.Fields{"id": id, "title": ticket.Title, "state": ticket.State}).Info("created ticket")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ticket)
}

func getAllTickets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tickets)
}

func getOneTicket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		// TODO: meaningful response message
	}

	for _, ticket := range tickets {
		if ticket.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ticket)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func updateTicket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		// TODO: meaningful response message
	}

	var newTicket Ticket
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		// TODO: meaningful response message
	}

	json.Unmarshal(reqBody, &newTicket)

	for _, ticket := range tickets {
		if ticket.ID == id {
			ticket.Title = newTicket.Title
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ticket)
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteTicket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		// TODO: meaningful response message
	}

	for i, ticket := range tickets {
		if ticket.ID == id {
			tickets = append(tickets[:i], tickets[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
