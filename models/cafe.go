package models

import (
	"github.com/gofrs/uuid"
	"googlemaps.github.io/maps"
)

type Cafe struct {
	ID      uuid.UUID   `json:"-"`
	PlaceID string      `json:"placeID"`
	Name    string      `json:"name"`
	LatLng  maps.LatLng `json:"location"`
}

type Cafes []Cafe

// LoadFromDetailsResult converts a Places API SearchResult into our Place model
func (c *Cafe) LoadFromDetailsResult(result maps.PlaceDetailsResult) {
	c.PlaceID = result.PlaceID
	c.Name = result.Name
	c.LatLng = result.Geometry.Location
}

// LoadFromSearchResults converts a Places API SearchResult into our Place model
func (c *Cafes) LoadFromSearchResults(results []maps.PlacesSearchResult) {
	for _, result := range results {
		*c = append(*c, Cafe{
			PlaceID: result.PlaceID,
			Name:    result.Name,
			LatLng:  result.Geometry.Location,
		})
	}
}
