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
	getIp, err := GetIpForPlace() //input worked in api playground
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", getIp)
	//%+v is special syntax to print struct to the console

	getLatLon, err := GetLatLonFromIp(IP)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", getLatLon)

}
