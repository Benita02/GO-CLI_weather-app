package main

import (
	"fmt"
	"net/http"
	"time"
)

// defining a global variable of type http.Client from the http package
// check documentation on it
var Client http.Client

func main() {
	Client = http.Client{
		//handy shortcuts for calculating time in the time package
		Timeout: time.Second * 10,
	}
	getLatLng, err := GetLatLngForPlace("38 Upper Montagu Street, London W1H 1LJ, United Kingdom") //input worked in api playground
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", getLatLng) //%+v is special syntax to print struct to the console

	//testing api on postman
	//trying another api, currently looking for an api to take in a location string and change to an ip address
}
