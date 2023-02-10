package service

import (
	"fmt"
	"io/ioutil"
	"net/http"

	config "example.com/openweather/config"
)

const OPEN_WEATHER_API_URL = "https://api.openweathermap.org/data/2.5/weather"

func GetWeatherByCity(cityName string) (response string) {
	openWeatherConfig := config.MyOpenWeatherConfig
	url := OPEN_WEATHER_API_URL

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return
	}

	// add query parameters
	q := req.URL.Query()
	q.Add("q", cityName)
	q.Add("appid", openWeatherConfig.Key)
	req.URL.RawQuery = q.Encode()

	// add header
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return
	}
	fmt.Printf("client: response body: %s\n", resBody)
	return string(resBody)
}
