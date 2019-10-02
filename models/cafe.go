package models

import (
	"googlemaps.github.io/maps"
)

type Place struct {
	Name   string      `json:"name"`
	LatLng maps.LatLng `json:"location"`
}

func (p *Place) LoadFromResult(result maps.PlacesSearchResult) {
	p.Name = result.Name
	p.LatLng = result.Geometry.Location
}
