package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"errors"
	"net/http"
	"net/url"
)

// master function to geocode oneline address using https://geocoding.geo.census.gov/geocoder/locations/onelineaddress
func geocode(addr string) (*Coordinate, error) {
	addr = url.QueryEscape(addr)
	reqUrl := fmt.Sprintf("https://geocoding.geo.census.gov/geocoder/locations/onelineaddress?address=%s&benchmark=2020&format=json", addr)

	resp, err := http.Get(reqUrl)
	if err != nil {
		return nil, errors.New("Could not GET URL for geocoding")
	}
	defer resp.Body.Close()

	var respBytes []byte
	respBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Could not read geocoding response")
	}

	var respGeocode Geocoding;
	err = json.Unmarshal(respBytes, &respGeocode)
	if err != nil {
		return nil, errors.New("Could not unmarshal geocoding response")
	}

	if len(respGeocode.Result.Addressmatches) == 0 {
		return nil, errors.New("No address found while geocoding")
	}

	return &respGeocode.Result.Addressmatches[0].Coordinates, nil
}
