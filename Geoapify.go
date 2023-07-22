package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IpLatLon struct {
	Ip  string  `json:"ip"`
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

type Location struct {
	Lat float64 `json:"latitude"`
	Lon float64 `json:"longitude"`
}

type Result struct {
	Location Location `json:"location"`
	Ip       string   `json:"ip"`
}

func GetLatLonForPlace() (LatLon IpLatLon, err error) {
	url := fmt.Sprintf("https://api.geoapify.com/v1/ipinfo?&apiKey=%s", Geoapify)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return LatLon, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	//client.Do - It handles the entire lifecycle of the request,
	//including establishing a connection, sending the request, and receiving the response.
	if err != nil {
		fmt.Println("Error making request:", err)
		return LatLon, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return LatLon, err
		}

		var result Result
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println("Error decoding response body:", err)
			return LatLon, err
		}

		Latlon := IpLatLon{
			Lat: result.Location.Lat,
			Lon: result.Location.Lon,
		}
		return Latlon, nil
	} else {
		fmt.Println("Error:", resp.StatusCode, resp.Status)
		return LatLon, fmt.Errorf("API request failed, error: %d", resp.StatusCode)
	}
}
