package golenoid

import (
	"fmt"
	"math"
)

// FieldPoint is an interface for a point in 3D space with a position and a magnetic field.
type FieldPoint interface {
	// GetPolarCoordinates returns the polar coordinates of the point.
	GetPolarCoordinates() (r, phi, z float64)
	// GetPolarField returns the magnetic field in polar coordinates.
	GetPolarField() (Br, Bphi, Bz float64)
	// GetCartesianCoordinates returns the cartesian coordinates of the point.
	GetCartesianCoordinates() (x, y, z float64)
	// GetCartesianField returns the magnetic field in cartesian coordinates.
	GetCartesianField() (Bx, By, Bz float64)
	// SetFieldPolar sets the magnetic field from polar components.
	SetFieldPolar(Br, Bphi, Bz float64)
	// SetFieldCartesian sets the magnetic field from cartesian components.
	SetFieldCartesian(Bx, By, Bz float64)
	// Magnitude returns the magnitude of the magnetic field at this point.
	Magnitude() float64
}

var _ FieldPoint = (*CartesianPoint)(nil)

// CartesianPoint is a point in 3D space with a position and a magnetic field in cartesian coordinates (x,y,z)
type CartesianPoint struct {
	X  float64
	Y  float64
	Z  float64
	Bx float64
	By float64
	Bz float64
}

// NewCartesianPoint creates a new CartesianPoint from the given coordinates (x,y,z).
func NewCartesianPoint(x, y, z float64) *CartesianPoint {
	return &CartesianPoint{
		X: x,
		Y: y,
		Z: z,
	}
}

func (p *CartesianPoint) GetPolarCoordinates() (r, phi, z float64) {
	return CartesianToPolarCoords(p.X, p.Y, p.Z)
}

func (p *CartesianPoint) GetPolarField() (Br, Bphi, Bz float64) {
	return CartesianToPolarField(p.Bx, p.By, p.Bz, math.Atan2(p.Y, p.X))
}

func (p *CartesianPoint) GetCartesianCoordinates() (x, y, z float64) {
	return p.X, p.Y, p.Z
}

func (p *CartesianPoint) GetCartesianField() (Bx, By, Bz float64) {
	return p.Bx, p.By, p.Bz
}

func (p *CartesianPoint) SetFieldPolar(Br, Bphi, Bz float64) {
	p.Bx, p.By, p.Bz = PolarToCartesianField(Br, Bphi, Bz, math.Atan2(p.Y, p.X))
}

func (p *CartesianPoint) SetFieldCartesian(Bx, By, Bz float64) {
	p.Bx = Bx
	p.By = By
	p.Bz = Bz
}

func (p *CartesianPoint) Magnitude() float64 {
	return math.Sqrt(p.Bx*p.Bx + p.By*p.By + p.Bz*p.Bz)
}

func (p *CartesianPoint) String() string {
	return fmt.Sprintf("CartesianPoint{X: %f, Y: %f, Z: %f, Bx: %f, By: %f, Bz: %f}", p.X, p.Y, p.Z, p.Bx, p.By, p.Bz)
}

var _ FieldPoint = (*PolarPoint)(nil)

// PolarPoint is a point in 3D space with a position and a magnetic field in cylindrical polar coordinates (r,phi,z)
type PolarPoint struct {
	R    float64
	Phi  float64
	Z    float64
	Br   float64
	Bphi float64
	Bz   float64
}

// NewPolarPoint creates a new PolarPoint from the given coordinates (r,phi,z).
func NewPolarPoint(r, phi, z float64) *PolarPoint {
	return &PolarPoint{
		R:   r,
		Phi: phi,
		Z:   z,
	}
}

func (p *PolarPoint) GetPolarCoordinates() (r, phi, z float64) {
	return p.R, p.Phi, p.Z
}

func (p *PolarPoint) GetPolarField() (Br, Bphi, Bz float64) {
	return p.Br, p.Bphi, p.Bz
}

func (p *PolarPoint) GetCartesianCoordinates() (x, y, z float64) {
	return PolarToCartesianCoords(p.R, p.Phi, p.Z)
}

func (p *PolarPoint) GetCartesianField() (Bx, By, Bz float64) {
	return PolarToCartesianField(p.Br, p.Bphi, p.Bz, p.Phi)
}

func (p *PolarPoint) SetFieldPolar(Br, Bphi, Bz float64) {
	p.Br = Br
	p.Bphi = Bphi
	p.Bz = Bz
}

func (p *PolarPoint) SetFieldCartesian(Bx, By, Bz float64) {
	p.Bz, p.Bphi, p.Br = CartesianToPolarField(Bx, By, Bz, p.Phi)
}

func (p *PolarPoint) Magnitude() float64 {
	return math.Sqrt(p.Br*p.Br + p.Bphi*p.Bphi + p.Bz*p.Bz)
}

func (p *PolarPoint) String() string {
	return fmt.Sprintf("PolarPoint{R: %f, Phi: %f, Z: %f, Br: %f, Bphi: %f, Bz: %f}", p.R, p.Phi, p.Z, p.Br, p.Bphi, p.Bz)
}
