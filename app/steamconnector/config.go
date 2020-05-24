package steamconnector

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
)

type key int

const (
	mongoKey key = iota
)

// Config type contains application configs
type Config struct {
	APIKey        string        `yaml:"key"`
	MongoURI      string        `yaml:"mongo_uri"`
	CacheDuration time.Duration `yaml:"cache_duration"`
}

// AppConfig is the global variable used to access configs
var AppConfig Config

func init() {
	paths := []string{"./config.yml", "./bin/config.yml"}
	var handle []byte
	var err error
	for _, path := range paths {

		file, _ := filepath.Abs(path)
		handle, err = ioutil.ReadFile(file)
		if err == nil {
			break
		}
	}
	err = yaml.Unmarshal(handle, &AppConfig)
	if err != nil {
		log.Printf("No config loaded")
	}
}

// AddContext adds a Mongo client to the context of requests
func AddContext(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		client, err := mongo.Connect(r.Context(), options.Client().ApplyURI(AppConfig.MongoURI))
		if err != nil {
			log.Printf("Failed to connect to mong: %v", err)
		}
		defer client.Disconnect(r.Context())
		log.Println(r.Method, "-", r.RequestURI)
		//Add data to context
		ctx := context.WithValue(r.Context(), mongoKey, client)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
