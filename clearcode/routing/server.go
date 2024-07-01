package routing

import (
	handle "clear_code/handelers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting(handler handle.Handleit) {
	r := mux.NewRouter()
	r.HandleFunc("/signup", handler.Sign).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", r))

}
