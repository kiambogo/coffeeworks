package main

import (
	"net/http"

	"github.com/kiambogo/coffeeworks/models"
	"github.com/kiambogo/coffeeworks/support"
)

// CreateReview handles requests to create new reviews against a cafe
func CreateReview(w http.ResponseWriter, r *http.Request) {
	form := &models.CreateReviewForm{}
	err := support.UnmarshalJSON(r.Body, form)
	if err != nil {
		support.ReturnString(w, 400, "Invalid data passed for CreateReviewForm")
		return
	}

	validationErrors := form.Validate()
	if len(validationErrors) > 0 {
		support.ReturnJSON(w, 400, validationErrors)
		return
	}

	review := &models.Review{}
	review.LoadFromCreateForm(*form)

	if err := DB.Create(review).Error; err != nil {
		support.LogError(err, "CreateReview - saving review to db")
		return
	}
	support.ReturnPrettyJSON(w, 200, form)
}
