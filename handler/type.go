package handler

import (
	"encoding/json"
	"londonDaily/cache"
	"londonDaily/trains"
	"net/http"
)

type Handler struct {
	Trains trains.Trains
	caches map[string]cache.Cache
}

func New(train trains.Trains) *Handler {
	return &Handler{
		Trains: train,
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
