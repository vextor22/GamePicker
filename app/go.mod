module github.com/vextor22/GamePicker/m/v2

go 1.14

require (
	github.com/gorilla/mux v1.7.4
	github.com/vextor22/gamepicker/app/steamconnector v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.3.3
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace github.com/vextor22/gamepicker/app/steamconnector => ./steamconnector
