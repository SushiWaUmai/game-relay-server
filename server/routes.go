package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func heathcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, API!")
}

func createLobby(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "Not Implemented")
}

func getLobbies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "Not Implemented")
}

func websocket(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Printf("Trying to access lobby with id: %s...", id)

	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "Not Implemented")
}

func SetupRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heathcheck).Methods("GET")
	router.HandleFunc("/lobby", createLobby).Methods("POST")
	router.HandleFunc("/lobby", getLobbies).Methods("GET")
	router.HandleFunc("/lobby/{id}", websocket).Methods("GET")

	return router
}
