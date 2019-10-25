package main

import (
	"net/http"

	"github.com/gorilla/mux"
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

	// Asynchronously process the review
	go ProcessReview(review)

	support.ReturnPrettyJSON(w, 200, form)
}

// GetReview handles requests to retrieve a review by place ID
func GetReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	placeID, ok := vars["id"]
	if !ok {
		support.ReturnString(w, 400, "Place ID required")
	}

	reviews := &models.Reviews{}
	if err := DB.Where("place_id = ?", placeID).Find(&reviews).Error; err != nil {
		support.LogError(err, "GetReview (%v)", placeID)
		return
	}

	support.ReturnPrettyJSON(w, 200, reviews)
}
