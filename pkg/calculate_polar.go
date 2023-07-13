package golenoid

import (
	"math"

	"gonum.org/v1/gonum/mathext"
)

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

	Br = C * z / (2 * alpha * alpha * beta * r) * ((a*a+r*r+z*z)*mathext.CompleteE(ksq) - (alpha*alpha)*mathext.CompleteK(ksq))
	if math.IsNaN(Br) {
		// This handles the special case where the point at r=0.
		Br = 0
	}
	Bphi = 0
	Bz = C / (2 * alpha * alpha * beta) * ((a*a-r*r-z*z)*mathext.CompleteE(ksq) + alpha*alpha*mathext.CompleteK(ksq))
	return
}
