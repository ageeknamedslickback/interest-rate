package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCalculateInterestHandler(t *testing.T) {
	routeVariable := "1001"
	invalidRouteVariable := "invalid"
	negativeRouteVariable := "-1001"

	tt := []struct {
		name       string
		path       string
		statusCode int
	}{
		{
			name:       "happy case",
			path:       fmt.Sprintf("/interest/%s", routeVariable),
			statusCode: http.StatusOK,
		},
		{
			name:       "sad case - negative number",
			path:       fmt.Sprintf("/interest/%s", negativeRouteVariable),
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "sad case - not a str number",
			path:       fmt.Sprintf("/interest/%s", invalidRouteVariable),
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tt {
		req, err := http.NewRequest("GET", tc.path, nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()

		// To add the vars to the context,
		// we need to create a router through which we can pass the request.
		router := mux.NewRouter()
		router.HandleFunc("/interest/{balance}", CalculateInterestHandler)
		router.ServeHTTP(rr, req)
		router.ServeHTTP(rr, req)

		if rr.Code != tc.statusCode {
			t.Fatalf("expected status code %v, but got %v with body %v", tc.statusCode, rr.Code, rr.Body.String())
		}
	}
}
