package trains

import "londonDaily/cache"

type Trains struct {
	Location string
	username string
	password string
	cache    *cache.Cache
}

func New(username string, password string, location string) *Trains {
	return &Trains{
		username: username,
		password: password,
		Location: location,
		cache:    cache.New(60, nil),
	}
}

type TrainResponse struct {
	Location struct {
		Name    string `json:"name"`
		Crs     string `json:"crs"`
		Tiploc  string `json:"tiploc"`
		Country string `json:"country"`
		System  string `json:"system"`
	} `json:"location"`
	Filter   interface{} `json:"filter"`
	Services []struct {
		LocationDetail struct {
			RealtimeActivated   bool   `json:"realtimeActivated"`
			Tiploc              string `json:"tiploc"`
			Crs                 string `json:"crs"`
			Description         string `json:"description"`
			GbttBookedArrival   string `json:"gbttBookedArrival"`
			GbttBookedDeparture string `json:"gbttBookedDeparture"`
			Origin              []struct {
				Tiploc      string `json:"tiploc"`
				Description string `json:"description"`
				WorkingTime string `json:"workingTime"`
				PublicTime  string `json:"publicTime"`
			} `json:"origin"`
			Destination []struct {
				Tiploc      string `json:"tiploc"`
				Description string `json:"description"`
				WorkingTime string `json:"workingTime"`
				PublicTime  string `json:"publicTime"`
			} `json:"destination"`
			IsCall                  bool   `json:"isCall"`
			IsPublicCall            bool   `json:"isPublicCall"`
			RealtimeArrival         string `json:"realtimeArrival"`
			RealtimeArrivalActual   bool   `json:"realtimeArrivalActual"`
			RealtimeDeparture       string `json:"realtimeDeparture"`
			RealtimeDepartureActual bool   `json:"realtimeDepartureActual"`
			Platform                string `json:"platform"`
			PlatformConfirmed       bool   `json:"platformConfirmed"`
			PlatformChanged         bool   `json:"platformChanged"`
			ServiceLocation         string `json:"serviceLocation,omitempty"`
			DisplayAs               string `json:"displayAs"`
			CancelReasonCode        string `json:"cancelReasonCode,omitempty"`
			CancelReasonShortText   string `json:"cancelReasonShortText,omitempty"`
			CancelReasonLongText    string `json:"cancelReasonLongText,omitempty"`
		} `json:"locationDetail"`
		ServiceUid      string `json:"serviceUid"`
		RunDate         string `json:"runDate"`
		TrainIdentity   string `json:"trainIdentity"`
		RunningIdentity string `json:"runningIdentity"`
		AtocCode        string `json:"atocCode"`
		AtocName        string `json:"atocName"`
		ServiceType     string `json:"serviceType"`
		IsPassenger     bool   `json:"isPassenger"`
	} `json:"services"`
}
