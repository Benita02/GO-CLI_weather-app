package main

import "fmt"

type OpenWeatherCondition struct {
	Id          int
	Main        string
	Description string
	Icon        string
}


type OpenWeatherResponseCurrent struct{
	Dt int64
	Sunrise int64
	Sunset int64
	Temp float32
	Feels_like float32
	Pressure int
	Humidity int
	Dew_point float32
	Uvi float32
	Clouds int
	Visibility int
	Wind_speed float32
	Wind_deg int
	Weather []OpenWeatherCondition //an array that'll be defined later
}


type OpenWeatherResponseHourly struct{
	Dt int64
	Temp float32
	Feels_like float32
	Pressure int
	Humidity int
	Dew_point float32
	Clouds int
	Visibility int
	Wind_speed float32
	Wind_deg int
	Weather []OpenWeatherCondition
	

}

// func (w OpenWeatherResponseHourly) Output(units string) string {
// 	var unitAbbr string

// 	switch units {
// 	case UnitsMetric:
// 		unitAbbr = "C"
// 	case UnitsImperial:
// 		unitAbbr = "F"
// 	}


	
}
type OpenWeatherResponseOneCall struct{
	Current *OpenWeatherResponseCurrent 
	Hourly *[]OpenWeatherResponseHourly
	Daily *[]OpenWeatherResponseDaily
}

func getWeatherInfo(lat_lng LatLng, units string, period string) {weather OpenweatherResponseOneCall, err error}
