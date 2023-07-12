# Golenoid

A package to calculate magnetic fields from solenoids.

## About

This package aims enable the calculation of magnetic fields within (and to a lesser extent around), tightly wound solenoidal electromagnets. The primary goal is to be able to compute fields quickly and without the need for sophisticated field modelling software that uses computationally expensive methods such as Finite Element Analysis. The kinds of solenoids this package is aimed at are those usually found in MRI machines or particle accelerators.

The mathematics for these calculations can be found in [here](https://ntrs.nasa.gov/citations/20140002333).

The assumptions that are made for computations are:

- The solenoid is tightly wound. This means that there is no space between the windings.
- The solenoid windings are perfectly circular.
- The magnetic axis lies perfectly along the z-axis, through the center of the coil.
