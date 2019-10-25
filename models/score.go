package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Score struct {
	ID                      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	PlaceID                 string    `json:"placeID" gorm:"index;not null"`
	WifiSpeed               float32   `json:"wifiSpeed"`
	WifiSpeedWeight         int       `json:"-"`
	WifiRestrictions        float32   `json:"wifiRestrictions"`
	WifiRestrictionsWeight  int       `json:"-"`
	BeverageSelection       float32   `json:"beverageSelection"`
	BeverageSelectionWeight int       `json:"-"`
	BeverageQuality         float32   `json:"beverageQuality"`
	BeverageQualityWeight   int       `json:"beverageQualityWeight"`
	NoiseLevel              float32   `json:"noiseLevel"`
	NoiseLevelWeight        int       `json:"-"`
	FoodOptions             float32   `json:"foodOptionsLevel"`
	FoodOptionsWeight       int       `json:"-"`
	CreatedAt               time.Time `gorm:"index"`
	UpdatedAt               time.Time `gorm:"index"`
	DeletedAt               *time.Time
}

type Scores []Score

// FindForPlaceID returns the score for the given place ID
func (s *Score) LoadLatest(placeID string) error {
	return DB.Where("place_id = ?", placeID).Order("created_at desc").Limit(1).First(s).Error
}
