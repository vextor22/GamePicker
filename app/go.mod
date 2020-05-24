module github.com/vextor22/GamePicker/m/v2

go 1.14

require (
	github.com/gorilla/mux v1.7.4
	github.com/vextor22/gamepicker/app/steamconnector v0.0.0-00010101000000-000000000000
	go.mongodb.org/mongo-driver v1.3.3 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200506231410-2ff61e1afc86 // indirect
)

replace github.com/vextor22/gamepicker/app/steamconnector => ./steamconnector
