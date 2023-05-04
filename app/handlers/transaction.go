package hanlders

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/balance", createBalance).Methods(http.MethodPost)
	r.HandleFunc("/balance", getBalance).Methods(http.MethodGet)
}

func createBalance(w http.ResponseWriter, r *http.Request) {
	var filename any
	json.NewDecoder(r.Body).Decode(&filename)
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("aa"))
}
