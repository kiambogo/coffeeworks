package models

import (
	"fmt"
	"strings"

	"github.com/kiambogo/coffeeworks/support"
)

type Review struct {
	PlaceID string  `json:"place_id"`
	Badges  []Badge `json:"badges"`
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
}

type CreateReviewForm struct {
	PlaceID string          `json:"place_id"`
	Badges  map[string]bool `json:"badges"`
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
		errors["badges"] = []string{
			fmt.Sprintf("at least one valid badge is required per review. Acceptable badges: [ %v ]", strings.Join(BadgeNames, ", ")),
		}
	}

	return errors
}
