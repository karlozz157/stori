package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	hanlders "github.com/karlozz157/stori/app/handlers"
)

func main() {
	r := mux.NewRouter()

	transactionHandler := hanlders.NewTransactionHandler()
	transactionHandler.RegisterRoutes(r)

	if err := http.ListenAndServe(os.Getenv("SERVER_PORT"), r); err != nil {
		log.Printf("Error running server %s\n", err.Error())
	}
}
