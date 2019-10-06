package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateReviewHandlerSuccess(t *testing.T) {
	body := `{ "place_id":"12345",
               "badges": {"wifi":true, "seating":true, "service":false},
               "wifi_speed": 0,
               "wifi_restrictions": 1,
               "beverage_selection": 2,
               "beverage_quality": 3,
               "noise_level": 4,
               "food_options_level": 2 }`

	req, _ := http.NewRequest("POST", "/api/reviews", strings.NewReader(body))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateReview)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)
}

func TestCreateReviewHandlerFailure(t *testing.T) {
	type testScenario struct {
		Name                string
		Body                string
		ExpectedResponseMsg string
	}

	testScenarios := []testScenario{
		testScenario{
			Name:                "No place_id",
			Body:                `{"badges":{"wifi":true, "seating":true, "service":false}}`,
			ExpectedResponseMsg: `{"place_id":["place_id must be specified"]}`,
		},
		testScenario{
			Name:                "No badges",
			Body:                `{"place_id":"12345"}`,
			ExpectedResponseMsg: `{"badges":["at least one valid badge is required per review. Acceptable badges: [ beverages, food, outlets, seating, service, wifi ]"]}`,
		},
		testScenario{
			Name:                "One invalid badge",
			Body:                `{"place_id":"12345", "badges":{"puppers":true}}`,
			ExpectedResponseMsg: `{"badges":["at least one valid badge is required per review. Acceptable badges: [ beverages, food, outlets, seating, service, wifi ]"]}`,
		},
		testScenario{
			Name:                "Wifi speed rating < 1",
			Body:                `{"place_id":"12345", "badges":{"wifi":true}, "wifi_speed":-1}`,
			ExpectedResponseMsg: `{"wifi_speed":["wifi_speed, when specified, must be an integer between 0 and 4"]}`,
		},
		testScenario{
			Name:                "Wifi speed rating > 4",
			Body:                `{"place_id":"12345", "badges":{"wifi":true}, "wifi_speed":5}`,
			ExpectedResponseMsg: `{"wifi_speed":["wifi_speed, when specified, must be an integer between 0 and 4"]}`,
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
