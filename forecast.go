package main

import (
	"fmt"
	"errors"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func WeatherPoints(coordinate *Coordinate) (string, error) {
	reqUrl := fmt.Sprintf("https://api.weather.gov/points/%f,%f", coordinate.X, coordinate.Y)
	resp, err := http.Get(reqUrl)
	if err != nil {
		return "", errors.New("Could not GET URL for weather point")
	}
	defer resp.Body.Close()

	var respBytes []byte
	respBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Could not read weather points response")
	}

	var respWeatherPoint WeatherPoint
	err = json.Unmarshal(respBytes, &respWeatherPoint)
	if err != nil {
		return "", errors.New("Could not unmarshal weather points response")
	}
	return respWeatherPoint.Properties.Forecastgriddata, nil
}

func Forecasts(coordinates *Coordinate) (*ForecastProperties, error) {
	baseUrl, err := WeatherPoints(coordinates)
	if err != nil {
		return nil, err
	}

	forecastUrl := baseUrl + "/forecast"
	var resp *http.Response
	resp, err = http.Get(forecastUrl)
	if err != nil {
		return nil, errors.New("Could not GET URL for forecast")
	}

	var respBytes []byte
	respBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read forecast response")
	}

	var respForecast ForecastResp
	err = json.Unmarshal(respBytes, &respForecast)
	if err != nil {
		return nil, errors.New("Could not unmarshal forecast response")
	}

	return &respForecast.Properties, nil
}


