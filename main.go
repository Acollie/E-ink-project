package main

import (
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"londonDaily/handler"
	"londonDaily/trains"
	"net/http"
)

type Config struct {
	TrainUsername string `envconfig:"TRAIN_USERNAME" required:"true"`
	TrainPassword string `envconfig:"TRAIN_PASSWORD" required:"true"`
	TrainLocation string `envconfig:"TRAIN_LOCATION" required:"true"`
}

func main() {
	var config Config
	envconfig.MustProcess("", &config)

	trainHandler := trains.New(config.TrainUsername, config.TrainPassword, config.TrainLocation)

	h := handler.New(*trainHandler)
	r := mux.NewRouter()
	r.HandleFunc("/trains", h.GetTrains).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
