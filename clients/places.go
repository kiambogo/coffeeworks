package clients

import (
	"context"
	"log"

	"googlemaps.github.io/maps"

	"github.com/kiambogo/coffeeworks/models"
)

type PlacesIface interface {
	FindPlacesNearArea(latLng maps.LatLng, radius int) (models.Cafes, error)
	GetPlaceDetails(placeID string) (models.Cafe, error)
}

func InitializePlacesClient(apiKey string) PlacesIface {
	c, err := maps.NewClient(maps.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	return &PlacesClient{c}
}

type PlacesClient struct {
	*maps.Client
}

// FindPlacesNearArea searches for places near a particular point
func (p *PlacesClient) FindPlacesNearArea(latLng maps.LatLng, radius int) (models.Cafes, error) {
	searchRequest := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: latLng.Lat,
			Lng: latLng.Lng,
		},
		Radius:  uint(radius),
		Keyword: "coffee",
	}

	resp, err := p.NearbySearch(context.Background(), searchRequest)
	if err != nil {
		log.Printf("ERROR: Invoking NearbySearch - %v", err.Error())
		return nil, err
	}

	cafes := &models.Cafes{}
	cafes.LoadFromSearchResults(resp.Results)

	return *cafes, nil
}

// GetPlaceDetails searches for places near a particular point
func (p *PlacesClient) GetPlaceDetails(placeID string) (models.Cafe, error) {
	request := &maps.PlaceDetailsRequest{PlaceID: placeID}

	resp, err := p.PlaceDetails(context.Background(), request)
	if err != nil {
		log.Printf("ERROR: Invoking PlaceDetails - %v", err.Error())
		return models.Cafe{}, err
	}

	cafe := &models.Cafe{}
	cafe.LoadFromDetailsResult(resp)

	return *cafe, nil
}

// MockPlacesClient used for testing
type MockPlacesClient struct{}

// FindPlacesNearArea is a mocked method, returning some dummy data
func (m *MockPlacesClient) FindPlacesNearArea(latLng maps.LatLng, radius int) (models.Cafes, error) {
	log.Printf("Mock: Finding places near %v, radius %v", latLng, radius)

	places := models.Cafes{
		models.Cafe{Name: "Joe's Coffee"},
		models.Cafe{Name: "Blenzzzz"},
	}
	return places, nil
}

// GetPlaceDetails searches for places near a particular point
func (p *MockPlacesClient) GetPlaceDetails(placeID string) (models.Cafe, error) {
	log.Printf("Mock: Retreiving details for cafe %v", placeID)

	return models.Cafe{Name: "Joe's Coffee"}, nil
}
