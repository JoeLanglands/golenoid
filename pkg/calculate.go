package golenoid

import (
	"math"

	"gonum.org/v1/gonum/mathext"
)

// This file contains the primitive field calculation procedures.

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

// CalculateFieldFromLoopPolar calculates the polar components of the magnetic field at point (r,phi,z) induced from a current loop.
//
// The coordinate (0, 0, 0) lies at the very centre of the current loop.
// In a cylindrically symmetric solenoid, the magnetic field is only a function of r and z and Bphi = 0.
// The input variables are:
//   - current: the current in the loop in amperes
//   - a: the radius of the current loop in metres
//   - r: the r coordinate of the point in metres
//   - z: the z coordinate of the point in metres
func CalculateFieldFromLoopPolar(current, a, r, z float64) (Br, Bphi, Bz float64) {
	C := calculateC(current)
	alpha := calculateAlpha(a, r, z)
	beta := calculateBeta(a, r, z)
	ksq := calculateKsquared(alpha, beta)

	E := mathext.CompleteE(ksq)
	K := alpha * alpha * mathext.CompleteK(ksq)

	if r == 0 {
		// This handles the special case where the point lies on the magnetic axis (r=0).
		// Otherwise Br ends up being NaN or +/-Inf.
		Br = 0
	} else {
		Br = C * z / (2 * alpha * alpha * beta * r) * ((a*a+r*r+z*z)*E - K)
	}
	Bphi = 0
	Bz = C / (2 * alpha * alpha * beta) * ((a*a-r*r-z*z)*E + K)
	return
}

// CalculateFieldFromLoopCartesian calculates the cartesian components of the magnetic field at point (x,y,z) induced from a current loop.
//
// The coordinate (0,0,0) lies at the very centre of the current loop.
// The input variables are:
//   - current: the current in the loop in amperes
//   - a: the radius of the current loop in metres
//   - x: the x coordinate of the point in metres
//   - y: the y coordinate of the point in metres
//   - z: the z coordinate of the point in metres
func CalculateFieldFromLoopCartesian(current, a, x, y, z float64) (Bx, By, Bz float64) {
	r := math.Sqrt(x*x + y*y)
	rho := math.Sqrt(x*x + y*y + z*z)
	C := calculateC(current)
	alpha := calculateAlpha(a, r, z)
	beta := calculateBeta(a, r, z)
	ksq := calculateKsquared(alpha, beta)

	E := mathext.CompleteE(ksq)
	K := alpha * alpha * mathext.CompleteK(ksq)

	if x == 0 && y == 0 {
		Bx = 0
		By = 0
	} else {
		Bx = C * x * z / (2 * alpha * alpha * beta * r * r) * ((a*a+rho*rho)*E - K)
		By = C * y * z / (2 * alpha * alpha * beta * r * r) * ((a*a+rho*rho)*E - K)
	}
	Bz = C / (2 * alpha * alpha * beta) * ((a*a-rho*rho)*E + K)
	return
}
