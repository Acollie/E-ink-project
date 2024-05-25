package weather

import (
	"encoding/json"
	"fmt"
	"londonDaily/cache"
	"net/http"
	"strconv"
)

type Weather struct {
	lat        float32
	long       float32
	locationID int
	apiKey     string
	cache      *cache.Cache
}

func New(apiKey string, lat float32, long float32, locationID int) Weather {
	return Weather{
		apiKey:     apiKey,
		lat:        lat,
		long:       long,
		locationID: locationID,
		cache:      cache.New(360, nil),
	}
}

type Condition string

const (
	openWeather = "http://api.openweathermap.org/data/2.5/forecast?id=%s&appid=%s"
)

type Forecast struct {
	Times       []int
	Temps       []float32
	WindSpeed   []float32
	RainChance  []float32
	WeatherType []Condition
}

func (h Weather) Fetch() (*Forecast, error) {
	if h.cache.Fetch() != nil {
		println("cache hit")
		return h.cache.Result.(*Forecast), nil
	}
	url := fmt.Sprintf(openWeather, strconv.Itoa(h.locationID), h.apiKey)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var weatherResponse OpenWeatherResponse
	err = json.NewDecoder(resp.Body).Decode(&weatherResponse)
	if err != nil {
		return nil, err
	}

	responseWeather := parse(weatherResponse)
	//Update the cache
	h.cache.Update(responseWeather)

	return responseWeather, nil
}

func parse(resp OpenWeatherResponse) *Forecast {

	var times []int
	var temps []float32
	var windSpeed []float32
	var rainChance []float32
	var weatherType []Condition

	for _, item := range resp.List {
		times = append(times, item.Dt)
		temps = append(temps, float32(item.Main.FeelsLike-273.15))
		windSpeed = append(windSpeed, float32(item.Wind.Speed))
		weatherType = append(weatherType, Condition(item.Weather[0].Description))
	}

	return &Forecast{
		Times:       times,
		Temps:       temps,
		WindSpeed:   windSpeed,
		RainChance:  rainChance,
		WeatherType: weatherType,
	}
}

type OpenWeatherResponse struct {
	Cod     string `json:"cod"`
	Message int    `json:"message"`
	Cnt     int    `json:"cnt"`
	List    []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Visibility int     `json:"visibility"`
		Pop        float64 `json:"pop"`
		Sys        struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
		Rain  struct {
			H float64 `json:"3h"`
		} `json:"rain,omitempty"`
	} `json:"list"`
	City struct {
		Id    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
		Timezone   int    `json:"timezone"`
		Sunrise    int    `json:"sunrise"`
		Sunset     int    `json:"sunset"`
	} `json:"city"`
}
