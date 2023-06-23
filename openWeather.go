package main

import "fmt"

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
	Rain struct {
		
	}

	
}
type OpenWeatherResponseOneCall struct{
	Current *OpenWeatherResponseCurrent 
	Hourly *[]OpenWeatherResponseHourly
	Daily *[]OpenWeatherResponseDaily
}

func getWeatherInfo(lat_lng LatLng, units string, period string) {weather OpenweatherResponseOneCall, err error}
