package golenoid

import "math"

// RadiansToDegrees converts radians to degrees.
func RadiansToDegrees(radians float64) float64 {
	return radians * 180 / math.Pi
}

// DegreesToRadians converts degrees to radians.
func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// PolarToCartesianCoords converts cylindrical polar coordinates to cartesian coordinates.
func PolarToCartesianCoords(r, phi, z float64) (float64, float64, float64) {
	x := r * math.Cos(phi)
	y := r * math.Sin(phi)
	return x, y, z
}

// CartesianToPolarCoords converts cartesian coordinates to cylindrical polar coordinates.
func CartesianToPolarCoords(x, y, z float64) (float64, float64, float64) {
	r := math.Sqrt(x*x + y*y)
	phi := math.Atan2(y, x)
	return r, phi, z
}

// PolarToCartesianField converts cylindrical polar magnetic field components to cartesian magnetic field components.
func PolarToCartesianField(Br, Bphi, Bz, phi float64) (float64, float64, float64) {
	Bx := Br*math.Cos(phi) - Bphi*math.Sin(phi)
	By := Br*math.Sin(phi) + Bphi*math.Cos(phi)
	return Bx, By, Bz
}

// CartesianToPolarField converts cartesian magnetic field components to cylindrical polar magnetic field components.
func CartesianToPolarField(Bx, By, Bz, phi float64) (float64, float64, float64) {
	Br := Bx*math.Cos(phi) + By*math.Sin(phi)
	Bphi := -Bx*math.Sin(phi) + By*math.Cos(phi)
	return Br, Bphi, Bz
}
