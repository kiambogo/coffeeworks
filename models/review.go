package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
	"github.com/kiambogo/coffeeworks/support"
)

type Review struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key;"`
	CreatedAt         time.Time `gorm:"index"`
	UpdatedAt         time.Time `gorm:"index"`
	DeletedAt         *time.Time
	PlaceID           string    `json:"placeID" gorm:"index"`
	Badges            []Badge   `json:"badges"`
	WifiSpeed         nulls.Int `json:"wifiSpeed"`
	WifiRestrictions  nulls.Int `json:"wifiRestrictions"`
	BeverageSelection nulls.Int `json:"beverageSelection"`
	BeverageQuality   nulls.Int `json:"beverageQuality"`
	NoiseLevel        nulls.Int `json:"noiseLevel"`
	FoodOptions       nulls.Int `json:"foodOptionsLevel"`
}

type Reviews []Review

func (r *Review) String() string {
	str, _ := support.PrettyPrintJSON(r)
	return str
}

// LoadFromCreateForm will create a Review models from the create form
func (r *Review) LoadFromCreateForm(form CreateReviewForm) {
	r.PlaceID = form.PlaceID

	for name, selected := range form.Badges {
		b, ok := Badges[name]
		if !ok {
			continue
		}
		b.Selected = selected
		r.Badges = append(r.Badges, b)
	}

	r.WifiSpeed = form.WifiSpeed
	r.WifiRestrictions = form.WifiRestrictions
	r.BeverageSelection = form.BeverageSelection
	r.BeverageQuality = form.BeverageQuality
	r.NoiseLevel = form.NoiseLevel
	r.FoodOptions = form.FoodOptions
}

type CreateReviewForm struct {
	PlaceID           string          `json:"placeID"`
	Badges            map[string]bool `json:"badges"`
	WifiSpeed         nulls.Int       `json:"wifiSpeed"`
	WifiRestrictions  nulls.Int       `json:"wifiRestrictions"`
	BeverageSelection nulls.Int       `json:"beverageSelection"`
	BeverageQuality   nulls.Int       `json:"beverageQuality"`
	NoiseLevel        nulls.Int       `json:"noiseLevel"`
	FoodOptions       nulls.Int       `json:"foodOptionsLevel"`
}

// Validate checks to ensure that the form has the data it needs
func (f *CreateReviewForm) Validate() map[string][]string {
	errors := make(map[string][]string)

	if f.PlaceID == "" {
		errors["placeID"] = []string{"placeID must be specified"}
	}

	validBadges := make(map[string]bool)
	for name, s := range f.Badges {
		if _, ok := Badges[name]; !ok {
			continue
		} else {
			validBadges[name] = s
		}
	}

	if len(validBadges) == 0 {
		errors["badges"] = []string{fmt.Sprintf("at least one valid badge is required per review. Acceptable badges: [ %v ]", strings.Join(BadgeNames, ", "))}
	}

	validate04Rating(errors, f.NoiseLevel, "noiseLevel")
	validate04Rating(errors, f.WifiRestrictions, "wifiRestrictions")
	validate04Rating(errors, f.WifiSpeed, "wifiSpeed")
	validate04Rating(errors, f.BeverageQuality, "beverageQuality")
	validate04Rating(errors, f.BeverageSelection, "beverageSelection")

	validate02Rating(errors, f.FoodOptions, "foodOptionsLevel")

	return errors
}

func validate04Rating(errors map[string][]string, value nulls.Int, fieldName string) {
	if !value.Valid {
		return
	}
	if value.Int < 0 || value.Int > 4 {
		errors[fieldName] = []string{fmt.Sprintf("%v, when specified, must be an integer between 0 and 4", fieldName)}
	}
}

func validate02Rating(errors map[string][]string, value nulls.Int, fieldName string) {
	if !value.Valid {
		return
	}
	if value.Int < 0 || value.Int > 2 {
		errors[fieldName] = []string{fmt.Sprintf("%v, when specified, must be an integer between 0 and 2", fieldName)}
	}
}
