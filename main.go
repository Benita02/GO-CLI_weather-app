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
	getLatLng, err := GetLatLngForPlace("129.18.0.0") //input worked in api playground
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", getLatLng) //%+v is special syntax to print struct to the console

	//testing api on postman

}
