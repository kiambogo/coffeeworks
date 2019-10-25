package main

import (
	"testing"

	"github.com/kiambogo/coffeeworks/models"
	"github.com/kiambogo/coffeeworks/testsupport"
	"github.com/stretchr/testify/assert"
)

func TestProcessReview(t *testing.T) {
	setupTest()

	// Insert a preliminary score
	score := testsupport.ValidScore()
	models.DB.Create(&score)

	review := testsupport.ValidReview()

	err := ProcessReview(&review)
	assert.NoError(t, err)

	// Ensure an updated score is made
	count := 0
	models.DB.Model(&models.Score{}).Where("place_id = ?", score.PlaceID).Count(&count)
	assert.Equal(t, 2, count)

	// Ensure new score is accurate
	newScore := &models.Score{}
	models.DB.Where("place_id = ?", score.PlaceID).Order("created_at desc").Limit(1).First(newScore)

	assert.Equal(t, float32(3.487805), newScore.WifiSpeed)
	assert.Equal(t, 41, newScore.WifiSpeedWeight)
	assert.Equal(t, float32(2.7741935), newScore.WifiRestrictions)
	assert.Equal(t, 31, newScore.WifiRestrictionsWeight)
	assert.Equal(t, float32(0.29732147), newScore.BeverageSelection)
	assert.Equal(t, 112, newScore.BeverageSelectionWeight)
	assert.Equal(t, float32(2.2), newScore.BeverageQuality)
	assert.Equal(t, 1, newScore.BeverageQualityWeight)
	assert.Equal(t, float32(4), newScore.NoiseLevel)
	assert.Equal(t, 13, newScore.NoiseLevelWeight)
	assert.Equal(t, float32(3.7813952), newScore.FoodOptions)
	assert.Equal(t, 43, newScore.FoodOptionsWeight)
}
