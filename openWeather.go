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

func (w OpenWeatherResponseCurrent) Output(units string) string {
	var unitAbbr string

	switch units {
	case UnitsMetric:
		unitAbbr = "C"
	case UnitsImperial:
		unitAbbr = "F"
	}

	return fmt.Sprintf("Current: %g%s | Humidity: %d%% | %s\n",
		w.Temp,
		unitAbbr,
		w.Humidity,
		w.Weather[0].Description,sssssssssssssssssssssssssssss
	)
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

func (w OpenWeatherResponseHourly) Output(units string) string {
	var unitAbbr string

	switch units {
	case UnitsMetric:
		unitAbbr = "C"
	case UnitsImperial:
		unitAbbr = "F"
	}

	t := time.Unix(w.Dt, 0)
	return fmt.Sprintf("%-9s %2d/%2d %02d:00   %5.2f%s | Humidity: %d%% | %s\n",
		t.Weekday().String(),
		t.Month(),
		t.Day(),
		t.Hour(),
		w.Temp,
		unitAbbr,
		w.Humidity,
		w.Weather[0].Description,
	)
}
type OpenWeatherResponseOneCall struct{
	Current *OpenWeatherResponseCurrent 
	Hourly *[]OpenWeatherResponseHourly
	Daily *[]OpenWeatherResponseDaily
}

func getWeatherInfo(lat_lng LatLng, units string, period string) {weather OpenweatherResponseOneCall, err error}
