package main

import (
	"encoding/json"
	"fmt"
)

type LatLng struct {
	Lat float64
	Lon float64
}
type LatLngResult struct {
	Lat float64 `json: "latitude"`
	Lon float64 `json: "longitude"`
}

func GetLatLonFromIp(Ip Ip) (latLng LatLng, err error) {
	url := fmt.Sprintf("https://ipapi.co/%s/json/", IP)

	res, err := Client.Get(url)

	if err != nil {
		return latLng, err
	}

	defer res.Body.Close()

	var latAndLng LatLngResult
	err = json.NewDecoder(res.Body).Decode(&latAndLng)

	if err != nil {
		return latLng, err
	}

	LatAndLng := LatLng{
		Lat: latAndLng.Lat,
		Lon: latAndLng.Lon,
	}

	return LatAndLng, nil

}
