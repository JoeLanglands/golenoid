package golenoid

import (
	"testing"
)

func TestCalculateFieldFromLoopPolar(t *testing.T) {
	// TODO @JoeLanglands figure out a good tolerance, write more tests and change the expectedBz
	// in the only test case
	tolerance := 0.001

	tt :=
		[]struct {
			name       string
			current    float64
			radius     float64
			r          float64
			z          float64
			expectedBz float64
			expectedBr float64
		}{
			{name: "origin", current: 1, radius: 1, r: 0, z: 0, expectedBz: 0.000001, expectedBr: 0},
		}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actualBr, _, actualBz := CalculateFieldFromLoopPolar(tc.current, tc.radius, tc.r, tc.z)
			if !approxEqual(actualBr, tc.expectedBr, tolerance) {
				t.Errorf("expected %f, got %f", tc.expectedBz, actualBz)
			}
			if !approxEqual(actualBz, tc.expectedBz, tolerance) {
				t.Errorf("expected %f, got %f", tc.expectedBr, actualBr)
			}
		})
	}
}

func approxEqual(a, b, tolerance float64) bool {
	return (a-b) < tolerance && (b-a) < tolerance
}
