package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LatLng struct {
	Lat float64
	Lng float64
}

//	func (l LocationResponse) ToLatLng() LatLng {
//		return l.LocationRes.Location
//	}

type GeocodeResult struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

func GetLatLngForPlace(city, country string) (lat_lng LatLng, err error) { //how will i know when to do error handling

	url := fmt.Sprintf("https://api.api-ninjas.com/v1/geocoding?city=%s&country=%s", city, country)

	//since go does not have it's default http clients does not have a time-out for request,
	// we're going to define out own custom http client with a time-out
	//this customized timer is done in main.go

	res, err := Client.Get(url) //making GET request to API

	//when you make a http request using http.Client as done above, it returns an http.Response object
	//hence we can use the various fields and methods under the http.Response struct e.g res.StatusCode -
	// .StatusCode is field under the http.Response struct
	if err != nil {
		return lat_lng, err
	}
	//close folder(response)
	defer res.Body.Close()
	//read doc

	var geocode GeocodeResult
	err = json.NewDecoder(res.Body).Decode(&geocode)

	//The json.NewDecoder function takes an input source,
	//which implements the io.Reader interface, as its parameter.
	//check the documentation for all these functions

	if err != nil {
		return lat_lng, err
	}

	if res.StatusCode != http.StatusOK { //|| len(geocode) < 1 {
		return lat_lng, fmt.Errorf("API request failed, error: %d", res.StatusCode)
	}

	latLng := LatLng{
		Lat: geocode.Lat,
		Lng: geocode.Lon,
	}

	return latLng, err
}
