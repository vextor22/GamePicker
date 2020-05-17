package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vextor22/gamepicker/app/steamconnector"
)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	steamconnector.RegisterSteamEndpoints(r)

	log.Printf("Steam Connector Service started")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App does not have this endpoint")
}
