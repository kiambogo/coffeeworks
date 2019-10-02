package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kiambogo/coffeeworks/testsupport"
)

func TestCafeLoadFromResult(t *testing.T) {
	result := testsupport.ValidPlacesSearchResult()
	place := &Place{}
	place.LoadFromResult(result)

	assert.Equal(t, place.PlaceID, result.ID)
	assert.Equal(t, place.Name, result.Name)
	assert.Equal(t, place.LatLng, result.Geometry.Location)
}
