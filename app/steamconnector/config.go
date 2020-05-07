package steamconnector

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Config struct {
	ApiKey string `yaml:"key"`
}

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
		log.Printf("Ahh, ouch")
	}
	log.Printf("Key is: %v", AppConfig)
}
