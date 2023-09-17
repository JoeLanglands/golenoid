package main

import (
	"fmt"
	"math"
	"time"

	golenoid "github.com/JoeLanglands/golenoid/pkg"
)

func main() {
	solenoid := golenoid.NewSolenoid(0.25, 0.28, 1.31, 200, 0, 768, 64)

	rMin := 0.
	rMax := 0.24
	phiMin := 0.
	phiMax := 2 * math.Pi
	zMin := -1.0
	zMax := 1.0
	nr := 6
	nphi := 6
	nz := 2000

	// field := golenoid.NewFieldGridPolar(rMin, rMax, phiMin, phiMax, zMin, zMax, nr, nphi, nz)

	// fmt.Printf("Calculating entire field with %d points\n", len(field.Points))
	// t := time.Now()
	// solenoid.CalculateFullField(field)
	// fmt.Printf("Took %s\n", time.Since(t))

	// nWorkers := runtime.NumCPU()

	for nWorkers := 1; nWorkers <= 24; nWorkers++ {
		fmt.Printf("Calculating field using %d workers... ", nWorkers)
		t := time.Now()
		solenoid.CalculateFieldWithWorkers(rMin, rMax, phiMin, phiMax, zMin, zMax, nr, nphi, nz, nWorkers)
		fmt.Printf("Took %s\n", time.Since(t))
	}

}
