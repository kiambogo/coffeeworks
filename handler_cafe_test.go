package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/kiambogo/coffeeworks/models"
	"github.com/kiambogo/coffeeworks/testsupport"
	"github.com/stretchr/testify/assert"
)

func TestListCafesHandler(t *testing.T) {
	setupTest()

	type testScenario struct {
		Name                 string
		URL                  string
		ExpectedResponseCode int
	}

	testScenarios := []testScenario{
		testScenario{
			Name:                 "No query params",
			URL:                  "/api/cafes",
			ExpectedResponseCode: 400,
		},
		testScenario{
			Name:                 "With lat and lng",
			URL:                  "/api/cafes?lat=40.00&lng=-120.00",
			ExpectedResponseCode: 200,
		},
		testScenario{
			Name:                 "With lat and lng and radius",
			URL:                  "/api/cafes?lat=40.00&lng=-120.00&radius=500",
			ExpectedResponseCode: 200,
		},
	}

	for _, ts := range testScenarios {
		req, _ := http.NewRequest("GET", ts.URL, nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(ListCafes)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != ts.ExpectedResponseCode {
			t.Errorf("(Test %v) ListCafes returned code %v when %v was expectedhandler", ts.Name, status, ts.ExpectedResponseCode)
		}
	}
}

func TestGetCafeHandler(t *testing.T) {
	setupTest()

	req, _ := http.NewRequest("GET", "/api/cafes/12345", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "12345"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCafe)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 200 {
		t.Errorf("GetCafe returned code %v when 200 was expectedhandler", status)
	}

	assert.Contains(t, rr.Body.String(), "Joe's Coffee")
	assert.Contains(t, rr.Body.String(), `"score": null`)
}

func TestGetCafeHandlerScore(t *testing.T) {
	setupTest()

	score := testsupport.ValidScore()
	models.DB.Create(&score)

	req, _ := http.NewRequest("GET", "/api/cafes/12345", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "12345"})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCafe)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != 200 {
		t.Errorf("GetCafe returned code %v when 200 was expectedhandler", status)
	}

	assert.Contains(t, rr.Body.String(), "Joe's Coffee")
	assert.Contains(t, rr.Body.String(), `"wifiSpeed": 3.5`)
}
