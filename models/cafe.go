package models

import (
	"github.com/gofrs/uuid"
	"googlemaps.github.io/maps"
)

type Place struct {
	ID      uuid.UUID   `json:"-"`
	PlaceID string      `json:"placeID"`
	Name    string      `json:"name"`
	LatLng  maps.LatLng `json:"location"`
}

// LoadFromResult converts a Places API SearchResult into our Place model
func (p *Place) LoadFromResult(result maps.PlacesSearchResult) {
	p.PlaceID = result.ID
	p.Name = result.Name
	p.LatLng = result.Geometry.Location
}
