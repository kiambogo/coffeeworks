package models

import (
	"github.com/gofrs/uuid"
	"github.com/kiambogo/coffeeworks/support"
	"googlemaps.github.io/maps"
)

type Cafe struct {
	ID      uuid.UUID   `json:"-"`
	PlaceID string      `json:"placeID"`
	Name    string      `json:"name"`
	LatLng  maps.LatLng `json:"location"`
}

type Cafes []Cafe

func (c *Cafe) String() string {
	str, _ := support.PrettyPrintJSON(c)
	return str
}

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
