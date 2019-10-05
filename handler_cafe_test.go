package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListCafesHandler(t *testing.T) {
	initPlacesClient()
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
	}

	for _, ts := range testScenarios {
		req, _ := http.NewRequest("GET", ts.URL, nil)

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(GetCafes)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != ts.ExpectedResponseCode {
			t.Errorf("(Test %v) ListCafes returned code %v when %v was expectedhandler", ts.Name, status, ts.ExpectedResponseCode)
		}
	}
}
