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
type GeocodeResponse struct {
	Lat float64
	Lng float64
}

type GeocodeResult struct {
	Result []GeocodeResponse
}

func getLatLngForPlace(place string) (lat_lng LatLng, err error) {
	url := fmt.Sprintf("https://us1.locationiq.com/v1/search?key=%s&q=%s&format=json",
		LocationiqApiKey,
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
		return lat_lng, errd
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return lat_lng,
			fmt.Errorf("API request failed, error: %d", res.StatusCode)

	}

	var geocode LatLng

	err = json.NewDecoder(res.Body).Decode(&geocode)

	//The json.NewDecoder function takes an input source,
	//which implements the io.Reader interface, as its parameter.
	//check the documentation for all these functions

	if err != nil {
		return lat_lng, err
	}

	lat_lng = LatLng{
		Lat: geocode.Lat,
		Lng: geocode.Lng,
	}

	return lat_lng, err
	//There's no error but our function returns two values
	//Can be geocode.Results[0].Location
}
