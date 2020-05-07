package steamconnector

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func RegisterSteamEndpoints(r *mux.Router) {
	r.HandleFunc("/user/{id}", userLookup).Methods("GET")
}

func userLookup(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}
	params := mux.Vars(r)
	userId := params["id"]
	log.Printf("Request: %v", userId)

	req, err := http.NewRequest("GET", "http://api.steampowered.com/IPlayerService/GetOwnedGames/v1", nil)
	if err != nil {
		log.Printf("Error generating request string: %v", err)
		return
	}
	q := req.URL.Query()
	q.Add("key", AppConfig.ApiKey)
	q.Add("steamid", userId)
	req.URL.RawQuery = q.Encode()

	log.Printf("Request URL: %v", req.URL.String())

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error doing request: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp_body, _ := ioutil.ReadAll(resp.Body)
	w.Write(resp_body)
}
