package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vextor22/gamepicker/app/steamconnector"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "App does not have this endpoint")
}
