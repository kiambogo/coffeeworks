package testsupport

import "googlemaps.github.io/maps"

// ValidPlacesSearchResult returns a gmaps PlacesSearchResult for testing purposes
func ValidPlacesSearchResults() []maps.PlacesSearchResult {
	return []maps.PlacesSearchResult{
		maps.PlacesSearchResult{
			PlaceID: RandomString(40),
			Name:    "Bob Loblaw's Cafe",
			Geometry: maps.AddressGeometry{
				Location: maps.LatLng{
					Lat: 49.281574,
					Lng: -123.110028,
				},
			},
		},
	}
}

// ValidPlaceDetailsResult returns a gmaps PlacesSearchResult for testing purposes
func ValidPlaceDetailsResult() maps.PlaceDetailsResult {
	return maps.PlaceDetailsResult{
		PlaceID: RandomString(40),
		Name:    "Bob Loblaw's Cafe",
		Geometry: maps.AddressGeometry{
			Location: maps.LatLng{
				Lat: 49.281574,
				Lng: -123.110028,
			},
		},
	}
}
