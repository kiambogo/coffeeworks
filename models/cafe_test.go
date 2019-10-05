package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kiambogo/coffeeworks/testsupport"
)

func TestCafeLoadFromSearchResult(t *testing.T) {
	results := testsupport.ValidPlacesSearchResults()
	cafes := &Cafes{}
	cafes.LoadFromSearchResults(results)

	assert.Equal(t, 1, len(results))
	assert.Equal(t, 1, len(*cafes))

	assert.Equal(t, (*cafes)[0].PlaceID, results[0].PlaceID)
	assert.Equal(t, (*cafes)[0].Name, results[0].Name)
	assert.Equal(t, (*cafes)[0].LatLng, results[0].Geometry.Location)
}

func TestCafeLoadFromDetailsResult(t *testing.T) {
	result := testsupport.ValidPlaceDetailsResult()
	cafe := &Cafe{}
	cafe.LoadFromDetailsResult(result)

	assert.Equal(t, cafe.PlaceID, result.PlaceID)
	assert.Equal(t, cafe.Name, result.Name)
	assert.Equal(t, cafe.LatLng, result.Geometry.Location)
}
