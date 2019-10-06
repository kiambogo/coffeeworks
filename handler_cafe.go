package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kiambogo/coffeeworks/support"
	"googlemaps.github.io/maps"
)

// GetCafe queries for a particular cafe by id
func GetCafe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cafeID, ok := vars["id"]
	if !ok {
		support.ReturnString(w, 400, "Cafe ID required")
	}

	cafe, err := PlacesClient.GetPlaceDetails(cafeID)
	if err != nil {
		support.PrintError(err)
		support.ReturnString(w, 500, "Something went wrong, yo")
		return
	}

	support.ReturnPrettyJSON(w, 200, cafe)
}

// ListCafes queries for cafes around a certain point
func ListCafes(w http.ResponseWriter, r *http.Request) {
	latLng, err := parseLatLng(r)
	if err != nil {
		support.ReturnString(w, 400, "Need numeric query params 'lat' and 'lng'")
		return
	}

	radiusStr := support.GetQueryParamDefault(r, "radius", "100")
	radius, err := strconv.Atoi(radiusStr)
	if err != nil {
		support.ReturnString(w, 400, "Radius, when specified, must be of numeric type")
		return
	}

	places, err := PlacesClient.FindPlacesNearArea(latLng, radius)
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
