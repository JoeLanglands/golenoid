package golenoid

import (
	"math"
	"testing"
)

var resultBx, resultBy, resultBz float64

func BenchmarkCalculateFieldConcurrentlyPolar(b *testing.B) {
	var bi, bj, bk float64
	solenoid := NewSolenoid(0.25, 0.28, 1.31, 200, 0, 768, 64)
	fpPolar := NewPolarPoint(0, 0, 0)

	for i := 0; i < b.N; i++ {
		bi, bj, bk = solenoid.CalculateFieldAtPoint(fpPolar)
	}
	resultBx, resultBy, resultBz = bi, bj, bk
}

func BenchmarkCalculateFieldConcurrentlyCartesian(b *testing.B) {
	var bi, bj, bk float64
	solenoid := NewSolenoid(0.25, 0.28, 1.31, 200, 0, 768, 64)
	fpCart := NewCartesianPoint(0, 0, 0)

	for i := 0; i < b.N; i++ {
		bi, bj, bk = solenoid.CalculateFieldAtPoint(fpCart)
	}
	resultBx, resultBy, resultBz = bi, bj, bk
}

func BenchmarkCalculateFieldSequentiallyPolar(b *testing.B) {
	var bi, bj, bk float64
	solenoid := NewSolenoid(0.25, 0.28, 1.31, 200, 0, 768, 64)
	fpPolar := NewPolarPoint(0, 0, 0)

	for i := 0; i < b.N; i++ {
		bi, bj, bk = solenoid.CalculateFieldAtPointSeq(fpPolar)
	}
	resultBx, resultBy, resultBz = bi, bj, bk
}

func BenchmarkCalculateFieldSequentiallyCartesian(b *testing.B) {
	var bi, bj, bk float64
	solenoid := NewSolenoid(0.25, 0.28, 1.31, 200, 0, 768, 64)
	fpCart := NewCartesianPoint(0, 0, 0)

	for i := 0; i < b.N; i++ {
		bi, bj, bk = solenoid.CalculateFieldAtPointSeq(fpCart)
	}
	resultBx, resultBy, resultBz = bi, bj, bk
}

func BenchmarkCalculateFullField(b *testing.B) {
	solenoid := NewSolenoid(0.25, 0.28, 1.31, 200, 0, 768, 64)

	rMin := 0.
	rMax := 0.24
	phiMin := 0.
	phiMax := 2 * math.Pi
	zMin := -1.0
	zMax := 1.0
	nr := 6
	nphi := 6
	nz := 2000
	field := NewFieldGridPolar(rMin, rMax, phiMin, phiMax, zMin, zMax, nr, nphi, nz)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		solenoid.CalculateFullField(field)
	}
}
