package trains

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Trains) Fetch() (TrainResponse, error) {
	if h.cache.Fetch() != nil {
		println("Cache hit")
		return h.cache.Result.(TrainResponse), nil
	}
	apiURL := fmt.Sprintf(fetchForStation, h.Location)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return TrainResponse{}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(h.username, h.password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return TrainResponse{}, err
	}
	defer resp.Body.Close()

	var trainResponse TrainResponse
	err = json.NewDecoder(resp.Body).Decode(&trainResponse)
	if err != nil {
		return TrainResponse{}, err
	}

	//Update the cache
	h.cache.Update(trainResponse)

	return trainResponse, nil
}
