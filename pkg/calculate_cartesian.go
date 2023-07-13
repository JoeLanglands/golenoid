package golenoid

import (
	"math"

	"gonum.org/v1/gonum/mathext"
)

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
	// TODO @JoeLanglands Check these are correct
	r := math.Sqrt(x*x + y*y)
	rho := math.Sqrt(x*x + y*y + z*z)
	C := calculateC(current)
	alpha := calculateAlpha(a, r, z)
	beta := calculateBeta(a, r, z)
	ksq := calculateKsquared(alpha, beta)

	Bx = C * x * z / (2 * alpha * alpha * beta * rho * rho) * ((a*a+r*r)*mathext.CompleteE(ksq) - (alpha*alpha)*mathext.CompleteK(ksq))
	By = (y / x) * Bx
	Bz = C / (2 * alpha * alpha * beta) * ((a*a+r*r)*mathext.CompleteE(ksq) - alpha*alpha*mathext.CompleteK(ksq))
	return
}
