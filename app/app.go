package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	routers := mux.NewRouter()
	routers.HandleFunc("/time", timeHandler).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", routers))
}
