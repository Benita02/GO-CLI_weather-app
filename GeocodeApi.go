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
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	label string
	name string
	Type string
	number int
	street string
	postal_code int
	confidence int
	region string
	region_code string
	administrative_area string
	neighbourhood string
	country string
	country_code string
	map_url string
}

type GeocodeResponse struct {
	Data struct {
	   Results []GeocodeResult  `json:"results"`
	} `json:"data"`
 }

func getLatLngForPlace(place string) (lat_lng LatLng, err error) {
	url := fmt.Sprintf("http://api.positionstack.com/v1/forward?access_key=%s&query=%s",
		PositionStackApiKey,
		place,
	)

	//since go does not have it's default http clients does not have a time-out for request,
	// we're going to define out own custom http client with a time-out
	//this customized timer is done in main.go

	res, err := Client.Get(url) //making request to API

	//when you make a http request using http.Client as done above, it returns an http.Response object
	//hence we can use the various fields and methods under the http.Response struct e.g resp.StatusCode -
	// .StatusCode is field under the http.Response struct
	if err != nil {
		return lat_lng, err
	}

	defer res.Body.Close()

	var geocode GeocodeResponse

	err = json.NewDecoder(res.Body).Decode(&geocode)

	//The json.NewDecoder function takes an input source,
	//which implements the io.Reader interface, as its parameter.
	//check the documentation for all these functions

	if err != nil {
		return lat_lng, err
	}

	if res.StatusCode != http.StatusOK {
		return lat_lng,
		fmt.Errorf("API request failed, error: %d", res.StatusCode)

	}
	latLng := LatLng {
		Lat: geocode.Data.Results[0].Latitude,
		Lng: geocode.Data.Results[0].Longitude,
	}

	return latLng, nil		
}