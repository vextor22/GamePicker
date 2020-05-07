package steamconnector

import (
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type Config struct {
	ApiKey string `yaml:"key"`
}

var AppConfig Config

func RegisterSteamEndpoints(r *mux.Router) {
	r.HandleFunc("/user/{id}", userLookup).Methods("GET")
}

func userLookup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello"))
}

func init() {
	file, _ := filepath.Abs("./config.yaml")
	handle, _ := ioutil.ReadFile(file)
	err := yaml.Unmarshal(handle, &AppConfig)
	if err != nil {
		log.Printf("Ahh, ouch")
	}
	log.Printf("Key is: %v", AppConfig)
}
