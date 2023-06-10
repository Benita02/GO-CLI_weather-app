package main

import (
	"fmt"
	"net/http"
	"time"
)

//defining a global variable of type http.Client from the http package
//check documentation on it
var Client http.Client

func main(){ 
	Client = http.Client{
		//handy shortcuts for calculating time in the time package
		Timeout := time.Second * 10
	}
	getLatLng := getLatLngForPlace()
}