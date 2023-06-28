package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ip struct {
	Ip string
}

type IpResult struct {
	Ip string
}

func GetIpForPlace() (IpAddress Ip, err error) { //how will i know when to do error handling
	//escPlace := url.QueryEscape(place)
	url := fmt.Sprintf("https://api.geoapify.com/v1/ipinfo?&apiKey=%s", Geoapify)
	//since go does not have it's default http clients does not have a time-out for request,
	// we're going to define out own custom http client with a time-out
	//this customized timer is done in main.go

	res, err := Client.Get(url) //making GET request to API

	//when you make a http request using http.Client as done above, it returns an http.Response object
	//hence we can use the various fields and methods under the http.Response struct e.g res.StatusCode -
	// .StatusCode is field under the http.Response struct
	if err != nil {
		return IpAddress, err
	}
	//close folder(response)
	defer res.Body.Close()
	//read doc

	var result IpResult
	err = json.NewDecoder(res.Body).Decode(&result)

	//The json.NewDecoder function takes an input source,
	//which implements the io.Reader interface, as its parameter.
	//check the documentation for all these functions

	if err != nil {
		return IpAddress, err
	}

	if res.StatusCode != http.StatusOK { //|| len(geocode) < 1 {
		return IpAddress, fmt.Errorf("API request failed, error: %d", res.StatusCode)
	}

	IP := Ip{
		Ip: result.Ip,
	}
	return IP, err
}
