package main

import (
	"github.com/gobuffalo/nulls"
	"github.com/jinzhu/gorm"
	"github.com/kiambogo/coffeeworks/models"
	"github.com/kiambogo/coffeeworks/support"
)

// ProcessReview is used to update the score for a cafe based on a new review
func ProcessReview(review *models.Review) error {
	// Pull latest score
	score := &models.Score{}
	err := score.LoadLatest(review.PlaceID)
	if err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return support.LogError(err, "ProcessReview (%v) - retrieving score", review.ID.String())
		}
	}

	newScore := &models.Score{}
	newScore.PlaceID = review.PlaceID

	// Update each metric
	newScore.WifiSpeed, newScore.WifiSpeedWeight = updateOptionalRating(score.WifiSpeed, review.WifiSpeed, score.WifiSpeedWeight)
	newScore.WifiRestrictions, newScore.WifiRestrictionsWeight = updateOptionalRating(score.WifiRestrictions, review.WifiRestrictions, score.WifiRestrictionsWeight)
	newScore.BeverageSelection, newScore.BeverageSelectionWeight = updateOptionalRating(score.BeverageSelection, review.BeverageSelection, score.BeverageSelectionWeight)
	newScore.BeverageQuality, newScore.BeverageQualityWeight = updateOptionalRating(score.BeverageQuality, review.BeverageQuality, score.BeverageQualityWeight)
	newScore.NoiseLevel, newScore.NoiseLevelWeight = updateOptionalRating(score.NoiseLevel, review.NoiseLevel, score.NoiseLevelWeight)
	newScore.FoodOptions, newScore.FoodOptionsWeight = updateOptionalRating(score.FoodOptions, review.FoodOptions, score.FoodOptionsWeight)

	if err = models.DB.Create(newScore).Error; err != nil {
		return support.LogError(err, "ProcessReview (%v) - saving new score", review.ID.String())
	}

	return nil
}

func updateOptionalRating(rating float32, newRating nulls.Int, oldWeight int) (float32, int) {
	if !newRating.Valid {
		return rating, oldWeight
	}

	newWeight := oldWeight + 1

	return ((rating * float32(oldWeight)) + float32(newRating.Int)) / float32(newWeight), newWeight
}
