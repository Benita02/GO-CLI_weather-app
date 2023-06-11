package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type LatLng struct {
	Lat float64 
	Lng float64
}

type GeocodeResult struct {
	Lat float64 `json: "lat"`
	Lon float64 `json: "lon"`  
}

type LocationiqGeocodeRes struct {
	Result []GeocodeResult  
}

func getLatLngForPlace(place string) (lat_lng LatLng, err error){
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
		return lat_lng, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return lat_lng, 
		fmt.Errorf("API request failed, error: %d", res.StatusCode )
		
	}
	 
	var geocode LocationiqGeocodeRes

	err = json.NewDecoder(res.Body).Decode(&geocode)

	//The json.NewDecoder function takes an input source, 
	//which implements the io.Reader interface, as its parameter. 
	//check the documentation for all these functions

	if err != nil {
		return lat_lng, err
	}

//Checking response status
	if resp.StatusCode != http.StatusOK || len(geocode.Result) < 1 {
		return lat_lng, resp.StatusCode, err
	}

	return 

	lat_lng.Lat = geocode.Lat
	lat_lng.Lng = geocode.Lon

	return lat_lng, err
//There's no error but our function returns two values
}[ m a s t e r   0 f 2 4 8 6 8 ]   E v e r y t h i n g   t o   r e c e i v e   l a t   a n d   l o n g .   r e s p o n s e s   i s   i n   p l a c e   b u t   c o d e   i s   h a v i n g   i s s u e s   w i t h   g l o b a l   v a r i a b l e s  
   2   f i l e s   c h a n g e d ,   2 7   i n s e r t i o n s ( + ) ,   1 1   d e l e t i o n s ( - )  
 