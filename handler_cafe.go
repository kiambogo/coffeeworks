package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kiambogo/coffeeworks/support"
	"googlemaps.github.io/maps"
)

// GetCafes queries for cafes around a certain point
func GetCafes(w http.ResponseWriter, r *http.Request) {
// ListCafes queries for cafes around a certain point
func ListCafes(w http.ResponseWriter, r *http.Request) {
	latLng, err := parseLatLng(r)
	if err != nil {
		support.ReturnString(w, 400, "Need numeric query params 'lat' and 'lng'")
		return
	}

	places, err := PlacesClient.FindPlacesNearArea(latLng)
	if err != nil {
		support.PrintError(err)
		support.ReturnString(w, 500, "Something went wrong, yo")
		return
	}

	support.ReturnPrettyJSON(w, 200, places)
}

func parseLatLng(r *http.Request) (maps.LatLng, error) {
	var lat, lng float64
	var err error
	var latLng maps.LatLng

	latParams := support.GetQueryParam(r, "lat")
	lngParams := support.GetQueryParam(r, "lng")

	if len(latParams) == 0 || len(lngParams) == 0 {
		return latLng, fmt.Errorf("'lat' or 'lng' query params were not present")
	}

	errors := []error{}
	lat, err = strconv.ParseFloat(latParams[0], 64)
	if err != nil {
		errors = append(errors, err)
	}
	lng, err = strconv.ParseFloat(lngParams[0], 64)
	if err != nil {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		log.Printf("%v", errors)
		return latLng, fmt.Errorf("'lat' or 'lng' query params were not numeric types")
	}

	latLng = maps.LatLng{
		Lat: lat,
		Lng: lng,
	}

	return latLng, nil
}
