// Package golenoid is a package for calculating the magnetic field of a solenoid.
//
// Currently the package is focused on solenoids comprised of tightly wound coils
// such as those used in particle accelerators or MRI machines.
// Calculations are done in cylindrical polar coordinates (r,phi,z) and assuming that
// the coils are perfectly circular with the magnetic axis perfectly aligned with the z-axis.
//
// The mathematics behind the computations are described in the following paper:
// https://ntrs.nasa.gov/citations/20140002333
//
// Future support for rotated solenoids could be added in the future.
package golenoid
