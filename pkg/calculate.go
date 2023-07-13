package golenoid

import "math"

const (
	mu0 = 4e-7 * math.Pi // mu0 is the permeability of free space.
)

// calculateC calculates the C parameter.
// The current must be in amperes.
func calculateC(current float64) float64 {
	return mu0 * current / math.Pi
}

// calculateAlpha calculates the alpha parameter.
//
// This parameter makes the field calculations easier to understand.
// The input variables are:
//   - a: the radius of the current loop in metres
//   - r: the r coordinate of the point in metres
//   - z: the z coordinate of the point in metres
func calculateAlpha(a, r, z float64) float64 {
	return math.Sqrt(a*a + r*r + z*z - 2*a*r)
}

// calculateBeta calculates the beta parameter.
//
// This parameter makes the field calculations easier to understand.
// The input variables are:
//   - a: the radius of the current loop in metres
//   - r: the r coordinate of the point in metres
//   - z: the z coordinate of the point in metres
func calculateBeta(a, r, z float64) float64 {
	return math.Sqrt(a*a + r*r + z*z + 2*a*r)
}

// calculateK calculates the k parameter squared.
//
// This parameter (squared) is the argument to all the complete elliptic integrals.
// It always appears as k^2, so we do not take the root.
func calculateKsquared(alpha, beta float64) float64 {
	return 1 - (alpha * alpha / (beta * beta))
}
