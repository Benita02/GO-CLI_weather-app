package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	WeatherPeriodCurrent  = "current"
	WeatherPeriodMinutely = "minutely"
	WeatherPeriodHourly   = "hourly"
	WeatherPeriodDaily    = "daily"
	UnitsImperial         = "imperial"
	UnitsMetric           = "metric"
)

// defining a global variable of type http.Client from the http package
// check documentation on it
var Client http.Client

func main() {
	Client = http.Client{
		//handy shortcuts for calculating time in the time package
		Timeout: time.Second * 10,
	}
	l, err := GetLatLonForPlace() //input worked in api playground
	if err != nil {
		panic(err)
	}

	getWeather, err := getWeatherInfo(l, UnitsImperial, WeatherPeriodHourly)

	fmt.Printf("%+v\n", getWeather.Hourly)
	//%+v is special syntax to print struct to the console
	if err != nil {
		panic(err)
	}

}
