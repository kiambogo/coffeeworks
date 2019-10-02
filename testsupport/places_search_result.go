package testsupport

import "googlemaps.github.io/maps"

// ValidPlacesSearchResult returns a gmaps PlacesSearchResult for testing purposes
func ValidPlacesSearchResult() maps.PlacesSearchResult {
	return maps.PlacesSearchResult{
		ID:   RandomString(40),
		Name: "Bob Loblaw's Cafe",
		Geometry: maps.AddressGeometry{
			Location: maps.LatLng{
				Lat: 49.281574,
				Lng: -123.110028,
			},
		},
	}
}
