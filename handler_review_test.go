package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReviewHandlerSuccess(t *testing.T) {
	setupTest()

	body := `{ "placeID":"12345",
               "badges": {"wifi":true, "seating":true, "service":false},
               "wifiSpeed": 0,
               "wifiRestrictions": 1,
               "beverageSelection": 2,
               "beverageQuality": 3,
               "noiseLevel": 4,
               "foodOptionsLevel": 2 }`

	req, _ := http.NewRequest("POST", "/api/reviews", strings.NewReader(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateReview)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	// Ensure review saved to db
	count := 0
	DB.Table("reviews").Count(&count)
	assert.Equal(t, 1, count)
}

func TestCreateReviewHandlerFailure(t *testing.T) {
	type testScenario struct {
		Name                string
		Body                string
		ExpectedResponseMsg string
	}

	testScenarios := []testScenario{
		testScenario{
			Name:                "No placeID",
			Body:                `{"badges":{"wifi":true, "seating":true, "service":false}}`,
			ExpectedResponseMsg: `{"placeID":["placeID must be specified"]}`,
		},
		testScenario{
			Name:                "No badges",
			Body:                `{"placeID":"12345"}`,
			ExpectedResponseMsg: `{"badges":["at least one valid badge is required per review. Acceptable badges: [ beverages, food, outlets, seating, service, wifi ]"]}`,
		},
		testScenario{
			Name:                "One invalid badge",
			Body:                `{"placeID":"12345", "badges":{"puppers":true}}`,
			ExpectedResponseMsg: `{"badges":["at least one valid badge is required per review. Acceptable badges: [ beverages, food, outlets, seating, service, wifi ]"]}`,
		},
		testScenario{
			Name:                "Wifi speed rating < 1",
			Body:                `{"placeID":"12345", "badges":{"wifi":true}, "wifiSpeed":-1}`,
			ExpectedResponseMsg: `{"wifiSpeed":["wifiSpeed, when specified, must be an integer between 0 and 4"]}`,
		},
		testScenario{
			Name:                "Wifi speed rating > 4",
			Body:                `{"placeID":"12345", "badges":{"wifi":true}, "wifiSpeed":5}`,
			ExpectedResponseMsg: `{"wifiSpeed":["wifiSpeed, when specified, must be an integer between 0 and 4"]}`,
		},
	}

	for _, ts := range testScenarios {
		req, _ := http.NewRequest("POST", "/api/reviews", strings.NewReader(ts.Body))

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(CreateReview)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, ts.ExpectedResponseMsg, rr.Body.String(), ts.Name)
	}
}
