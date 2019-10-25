package testsupport

import (
	"github.com/gofrs/uuid"
	"github.com/kiambogo/coffeeworks/models"
)

func ValidScore() models.Score {
	return models.Score{
		ID:                      uuid.Must(uuid.NewV4()),
		PlaceID:                 "sweet-bean-12345",
		WifiSpeed:               3.5,
		WifiSpeedWeight:         40,
		WifiRestrictions:        2.8,
		WifiRestrictionsWeight:  30,
		BeverageSelection:       0.3,
		BeverageSelectionWeight: 111,
		BeverageQuality:         2.2,
		BeverageQualityWeight:   1,
		NoiseLevel:              4.0,
		NoiseLevelWeight:        12,
		FoodOptions:             3.8,
		FoodOptionsWeight:       42,
	}
}
