package golenoid

import (
	"math"
	"testing"
)

func TestRadiansToDegrees(t *testing.T) {
	// Note the names don't use slash as divide because it screws with the names
	// in the vscode test explorer :(
	tt := []struct {
		name     string
		radians  float64
		expected float64
	}{
		{name: "0", radians: 0, expected: 0},
		{name: "pi", radians: math.Pi, expected: 180},
		{name: "2pi", radians: 2 * math.Pi, expected: 360},
		{name: "pi_over_2", radians: math.Pi / 2, expected: 90},
		{name: "pi_over_4", radians: math.Pi / 4, expected: 45},
		{name: "3pi_over_4", radians: 3 * math.Pi / 4, expected: 135},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := RadiansToDegrees(tc.radians)
			if actual != tc.expected {
				t.Errorf("expected %f, got %f", tc.expected, actual)
			}
		})
	}
}

func TestDegreesToRadians(t *testing.T) {
	tt := []struct {
		name     string
		degrees  float64
		expected float64
	}{
		{name: "0 degrees", degrees: 0, expected: 0},
		{name: "180 degrees", degrees: 180, expected: math.Pi},
		{name: "360 degrees", degrees: 360, expected: 2 * math.Pi},
		{name: "90 degrees", degrees: 90, expected: math.Pi / 2},
		{name: "45 degrees", degrees: 45, expected: math.Pi / 4},
		{name: "135 degrees", degrees: 135, expected: 3 * math.Pi / 4},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := DegreesToRadians(tc.degrees)
			if actual != tc.expected {
				t.Errorf("expected %f, got %f", tc.expected, actual)
			}
		})
	}
}
