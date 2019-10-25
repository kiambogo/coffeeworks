package testsupport

import (
	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
	"github.com/kiambogo/coffeeworks/models"
)

func ValidReview() models.Review {
	return models.Review{
		ID:                uuid.Must(uuid.NewV4()),
		PlaceID:           "sweet-bean-12345",
		WifiSpeed:         nulls.NewInt(3),
		WifiRestrictions:  nulls.NewInt(2),
		BeverageSelection: nulls.NewInt(0),
		NoiseLevel:        nulls.NewInt(4),
		FoodOptions:       nulls.NewInt(3),
	}
}
