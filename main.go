package main

import (
	"fmt"
	"net/http"
	"time"
)

//defining a global variable of type http.Client from the http package
//check documentation on it
var CLIENT http.Client

func main(){ 
	CLIENT = http.Client{
		//handy shortcuts for calculationg time in the time package
		Timeout := time.Second * 10
	}
	getLatLng := getLatLngForPlace()
}