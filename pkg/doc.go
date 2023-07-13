// Package golenoid is a package for calculating the magnetic field of a solenoid.
//
// Currently the package is focused on solenoids comprised of tightly wound coils
// such as those used in particle accelerators or MRI machines. This is due to the mathematics
// relying on simplified expressions that assume the cross-section of the conductor is negligible.
//
// Calculations are performed assuming that the magnetic axis of the solenoid is collinear
// with the z-axis of the coordinate system, and infact the z-axis runs through the exact centre
// of the coil.
//
// Although assumptions are made, the expressions used to calculate the magnetic field are
// valid in all space outside the conductor and are exact solutions that satisfy Maxwell's equations.
//
// All units must be in SI units! (metres, amperes, teslas, etc.)
//
// The mathematics behind the computations are described in the following paper:
// https://ntrs.nasa.gov/citations/20140002333
//
// Future support for rotated solenoids and various transformations is planned.
package golenoid
