package main

type LatLng struct {
	Lat float64
	Lng float64
}

func getLatLngForPlace(place string) (lat_lng LatLng, err error){
	url := fmt.Sprintf("https://us1.locationiq.com/v1/search?key=%s&q=%s", 
		geocodeApiKey,
		place,
	)

//since go does not have it's default http clients does not have a time-out for request,
// we're going to define out own custom http client with a time-out

}
