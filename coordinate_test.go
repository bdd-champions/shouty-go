package main

import (
	"math"
	"testing"
)

func TestCoordinateDistanceFrom(t *testing.T) {
	testCases := []struct {
		name     string
		coord1   *Coordinate
		coord2   *Coordinate
		expected float64
	}{
		{
			name:     "distance from itself",
			coord1:   NewCoordinate(0, 0),
			coord2:   NewCoordinate(0, 0),
			expected: 0,
		},
		{
			name:     "distance from another coordinate along X axis",
			coord1:   NewCoordinate(0, 0),
			coord2:   NewCoordinate(600, 0),
			expected: 600,
		},
		// Uncomment and update this test case to match your requirements
		// {
		// 	name:     "distance from another coordinate",
		// 	coord1:   NewCoordinate(0, 0),
		// 	coord2:   NewCoordinate(300, 400),
		// 	expected: 500,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.coord1.DistanceFrom(tc.coord2)
			if math.Abs(got-tc.expected) > 1e-9 {
				t.Errorf("Distance from %v to %v = %f; want %f", tc.coord1, tc.coord2, got, tc.expected)
			}
		})
	}
}
