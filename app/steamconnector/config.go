package steamconnector

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
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
		log.Printf("No config loaded")
	}
}
