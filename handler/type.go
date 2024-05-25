package handler

import (
	"encoding/json"
	"londonDaily/cache"
	"londonDaily/trains"
	"londonDaily/weather"
	"net/http"
)

type Handler struct {
	Trains  trains.Trains
	Weather weather.Weather
	caches  map[string]cache.Cache
}

func New(train trains.Trains, weather weather.Weather) *Handler {
	return &Handler{
		Trains:  train,
		Weather: weather,
	}
}

func (h *Handler) GetTrains(w http.ResponseWriter, r *http.Request) {
	timing, err := h.Trains.Fetch()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(timing)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetWeather(w http.ResponseWriter, r *http.Request) {
	forecast, err := h.Weather.Fetch()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(forecast)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Service is running"))
}
