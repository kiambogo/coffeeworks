package models

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/nulls"
	"github.com/kiambogo/coffeeworks/support"
)

type Review struct {
	PlaceID           string    `json:"place_id"`
	Badges            []Badge   `json:"badges"`
	WifiSpeed         nulls.Int `json:"wifi_speed"`
	WifiRestrictions  nulls.Int `json:"wifi_restrictions"`
	BeverageSelection nulls.Int `json:"beverage_selection"`
	BeverageQuality   nulls.Int `json:"beverage_quality"`
	NoiseLevel        nulls.Int `json:"noise_level"`
	FoodOptions       nulls.Int `json:"food_options_level"`
}

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
	PlaceID           string          `json:"place_id"`
	Badges            map[string]bool `json:"badges"`
	WifiSpeed         nulls.Int       `json:"wifi_speed"`
	WifiRestrictions  nulls.Int       `json:"wifi_restrictions"`
	BeverageSelection nulls.Int       `json:"beverage_selection"`
	BeverageQuality   nulls.Int       `json:"beverage_quality"`
	NoiseLevel        nulls.Int       `json:"noise_level"`
	FoodOptions       nulls.Int       `json:"food_options_level"`
}

// Validate checks to ensure that the form has the data it needs
func (f *CreateReviewForm) Validate() map[string][]string {
	errors := make(map[string][]string)

	if f.PlaceID == "" {
		errors["place_id"] = []string{"place_id must be specified"}
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

	validate04Rating(errors, f.NoiseLevel, "noise_level")
	validate04Rating(errors, f.WifiRestrictions, "wifi_restrictions")
	validate04Rating(errors, f.WifiSpeed, "wifi_speed")
	validate04Rating(errors, f.BeverageQuality, "beverage_quality")
	validate04Rating(errors, f.BeverageSelection, "beverage_selection")

	validate02Rating(errors, f.FoodOptions, "food_options_level")

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
