package main

import (
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"londonDaily/handler"
	"londonDaily/trains"
	"londonDaily/weather"
	"net/http"
)

type Config struct {
	TrainUsername   string  `envconfig:"TRAIN_USERNAME" required:"true"`
	TrainPassword   string  `envconfig:"TRAIN_PASSWORD" required:"true"`
	TrainLocation   string  `envconfig:"TRAIN_LOCATION" required:"true"`
	WeatherLocation int     `envconfig:"WEATHER_LOCATION" required:"true"`
	WeatherLong     float32 `envconfig:"WEATHER_LONG" required:"true"`
	WeatherLat      float32 `envconfig:"WEATHER_LAT" required:"true"`
	WeatherAPI      string  `envconfig:"WEATHER_API" required:"true"`
}

func main() {
	var config Config
	envconfig.MustProcess("", &config)

	trainHandler := trains.New(config.TrainUsername, config.TrainPassword, config.TrainLocation)
	weatherHandler := weather.New(config.WeatherAPI, config.WeatherLat, config.WeatherLong, config.WeatherLocation)

	h := handler.New(*trainHandler, weatherHandler)
	r := mux.NewRouter()
	r.HandleFunc("/trains", h.GetTrains).Methods(http.MethodGet)
	r.HandleFunc("/weather", h.GetWeather).Methods(http.MethodGet)
	r.HandleFunc("/", h.GetHealthCheck).Methods(http.MethodGet)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)

}
