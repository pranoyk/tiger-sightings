package service

import (
	"context"
	"testing"

	"github.com/pranoyk/tiger-sightings/model"
)

func TestIsDistantFromLastSighting(t *testing.T) {
	// Create a test tiger instance
	testTiger := &tiger{}

	// Define test cases
	testCases := []struct {
		name            string
		allowedDistance float64
		lastSighting    *model.TigerSightings
		currentSighting *model.TigerSightings
		expectedResult  bool
	}{
		{
			name:            "is not distant",
			allowedDistance: 10.0,
			lastSighting: &model.TigerSightings{
				Lat: 10.0,
				Lng: 20.0,
			},
			currentSighting: &model.TigerSightings{
				Lat: 10.0,
				Lng: 20.1,
			},
			expectedResult: false,
		},
		{
			name:            "is greter than expected distance",
			allowedDistance: 2.0,
			lastSighting: &model.TigerSightings{
				Lat: 10.0,
				Lng: 20.0,
			},
			currentSighting: &model.TigerSightings{
				Lat: 11.5,
				Lng: 21.5,
			},
			expectedResult: true,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := testTiger.isDistantFromLastSighting(context.Background(), tc.allowedDistance, tc.lastSighting, tc.currentSighting)

			if result != tc.expectedResult {
				t.Errorf("Expected %v, but got %v", tc.expectedResult, result)
			}
		})
	}
}
