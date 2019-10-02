package clients

import (
	"context"
	"log"

	"googlemaps.github.io/maps"

	"github.com/kiambogo/coffeeworks/models"
)

type PlacesIface interface {
	FindPlacesNearArea(latLng maps.LatLng) ([]models.Place, error)
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
func (p *PlacesClient) FindPlacesNearArea(latLng maps.LatLng) ([]models.Place, error) {
	searchRequest := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: latLng.Lat,
			Lng: latLng.Lng,
		},
		Radius:  100,
		Keyword: "coffee",
	}

	resp, err := p.NearbySearch(context.Background(), searchRequest)
	if err != nil {
		log.Printf("ERROR: Invoking NearbySearch - %v", err.Error())
		return nil, err
	}

	places := []models.Place{}
	for _, result := range resp.Results {
		place := &models.Place{}
		place.LoadFromResult(result)

		places = append(places, *place)
	}

	return places, nil
}

// MockPlacesClient used for testing
type MockPlacesClient struct{}

// FindPlacesNearArea is a mocked method, returning some dummy data
func (m *MockPlacesClient) FindPlacesNearArea(latLng maps.LatLng) ([]models.Place, error) {
	log.Printf("Mock: Finding places near %v", latLng)

	places := []models.Place{
		models.Place{Name: "Joe's Coffee"},
		models.Place{Name: "Blenzzzz"},
	}
	return places, nil
}
