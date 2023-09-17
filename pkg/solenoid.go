package golenoid

import (
	"fmt"
	"sync"
)

// Solenoid represents a tightly wound solenoid
type Solenoid struct {
	Rinner    float64 // Inner radius of solenoid
	Router    float64 // Outer radius of solenoid
	Length    float64 // Length of solenoid
	Current   float64 // Current in solenoid (in Amperes)
	Nturns    int     // Number of turns per layer in solenoid layer
	Nlayers   int     // Number of layers in solenoid
	CentrePos float64 // Position of centre of solenoid along z-axis
}

// NewSolenoid creates a new Solenoid with the given parameters.
func NewSolenoid(rInner, rOuter, length, current, centre float64, nTurns, nLayers int) *Solenoid {
	return &Solenoid{
		Rinner:    rInner,
		Router:    rOuter,
		Length:    length,
		Current:   current,
		Nturns:    nTurns,
		Nlayers:   nLayers,
		CentrePos: centre,
	}
}

// TODO @JoeLanglands figure out how to parallelise this better. The mutex way is slower probably because of the locking.
// Figure out a way to use workers and channels etc (read your book.)
// You also have to achieve a balance because you could be spawning to many goroutines.

func (s *Solenoid) CalculateFullField(field *Field) {
	var wg sync.WaitGroup
	for i, p := range field.Points {
		wg.Add(1)
		go func(fp FieldPoint, i int) {
			defer wg.Done()
			// r, phi, z := fp.GetPolarCoordinates()
			Bi, Bj, Bk := s.CalculateFieldAtPointSeq(fp)
			// fmt.Printf("Point %d: P = (%f, %f, %f); B = (%f, %f, %f)\n", i, r, phi, z, Bi, Bj, Bk)
			switch p := fp.(type) {
			case *CartesianPoint:
				p.SetFieldCartesian(Bi, Bj, Bk) // Bear in mind this will not actually work because we're passing in copies
			case *PolarPoint:
				p.SetFieldPolar(Bi, Bj, Bk)
			default:
				panic(fmt.Sprintf("Unsupported point type: %T", p))
			}
		}(p, i)
	}
	wg.Wait()
}

// CalculateFieldAtPoint calculates the magnetic field at the point fp induced by the solenoid.
func (s *Solenoid) CalculateFieldAtPoint(fp FieldPoint) (Bi, Bj, Bk float64) {
	zStart := s.CentrePos - s.Length/2
	height := s.Router - s.Rinner

	loopSeparation := s.Length / float64(s.Nturns)
	layerSeparation := height / float64(s.Nlayers)

	Bi, Bj, Bk = s.calculateFieldOverLayers(fp, layerSeparation, loopSeparation, zStart, s.Rinner+0.5*layerSeparation)

	return
}

// calculateFieldOverLayers calculates the field at p over all layers in a solenoid.
func (s *Solenoid) calculateFieldOverLayers(fp FieldPoint, layerSep, loopSep, zStart, firstLoopR float64) (Bi, Bj, Bk float64) {
	loopRadius := firstLoopR - layerSep
	// var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(s.Nlayers)
	for i := 0; i < s.Nlayers; i++ {
		loopRadius += layerSep
		// i := i
		go func(r float64) {
			// fmt.Printf("Calculating layer %d out of %d\n", i, s.Nlayers)
			defer wg.Done()
			bi, bj, bk := s.calculateFieldOverLoops(fp, loopSep, zStart, r)
			// mu.Lock()
			Bi += bi
			Bj += bj
			Bk += bk
			// mu.Unlock()
		}(loopRadius)
	}
	wg.Wait()
	return
}

// calculateFieldOverLoops calculates the field at p over all loops in a layer.
func (s *Solenoid) calculateFieldOverLoops(fp FieldPoint, loopSep, zStart, loopRadius float64) (Bi, Bj, Bk float64) {
	var newZ float64
	offset := zStart - loopSep
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(s.Nturns)
	for i := 0; i < s.Nturns; i++ {
		offset += loopSep
		// i := i
		switch p := fp.(type) {
		case *CartesianPoint:
			newZ = p.Z - offset
			go func(z float64) {
				// fmt.Printf("Calculating loop %d out of %d\n", i+1, s.Nturns)
				defer wg.Done()
				bx, by, bz := CalculateFieldFromLoopCartesian(s.Current, loopRadius, p.X, p.Y, z)
				mu.Lock()
				Bi += bx
				Bj += by
				Bk += bz
				mu.Unlock()
			}(newZ)
		case *PolarPoint:
			newZ = p.Z - offset
			go func(z float64) {
				defer wg.Done()
				br, bphi, bz := CalculateFieldFromLoopPolar(s.Current, loopRadius, p.R, z)
				mu.Lock()
				Bi += br
				Bj += bphi
				Bk += bz
				mu.Unlock()
			}(newZ)
		default:
			panic(fmt.Sprintf("Unsupported point type: %T", p))
		}
	}
	wg.Wait()
	return
}

func (s *Solenoid) CalculateFieldPoint(fp FieldPoint) FieldPoint {
	zStart := s.CentrePos - s.Length/2
	height := s.Router - s.Rinner

	loopSeparation := s.Length / float64(s.Nturns)
	layerSeparation := height / float64(s.Nlayers)

	Bi, Bj, Bk := s.calculateFieldOverLayersSeq(fp, layerSeparation, loopSeparation, zStart, s.Rinner+0.5*layerSeparation)
	switch p := fp.(type) {
	case *CartesianPoint:
		p.SetFieldCartesian(Bi, Bj, Bk) // Bear in mind this will not actually work because we're passing in copies
	case *PolarPoint:
		p.SetFieldPolar(Bi, Bj, Bk)
	default:
		panic(fmt.Sprintf("Unsupported point type: %T", p))
	}
	return fp
}

func (s *Solenoid) CalculateFieldAtPointSeq(fp FieldPoint) (Bi, Bj, Bk float64) {
	zStart := s.CentrePos - s.Length/2
	height := s.Router - s.Rinner

	loopSeparation := s.Length / float64(s.Nturns)
	layerSeparation := height / float64(s.Nlayers)

	Bi, Bj, Bk = s.calculateFieldOverLayersSeq(fp, layerSeparation, loopSeparation, zStart, s.Rinner+0.5*layerSeparation)

	return
}

func (s *Solenoid) calculateFieldOverLayersSeq(fp FieldPoint, layerSep, loopSep, zStart, firstLoopR float64) (Bi, Bj, Bk float64) {
	loopRadius := firstLoopR - layerSep

	for i := 0; i < s.Nlayers; i++ {
		loopRadius += layerSep
		// fmt.Printf("Calculating layer %d out of %d\n", i, s.Nlayers)
		bi, bj, bk := s.calculateFieldOverLoopsSeq(fp, loopSep, zStart, loopRadius)
		Bi += bi
		Bj += bj
		Bk += bk
	}
	return
}

func (s *Solenoid) calculateFieldOverLoopsSeq(fp FieldPoint, loopSep, zStart, loopRadius float64) (Bi, Bj, Bk float64) {
	var newZ float64
	offset := zStart - loopSep

	for i := 0; i < s.Nturns; i++ {
		offset += loopSep
		// fmt.Println("Calculating loop", i, "out of", s.Nturns)
		switch p := fp.(type) {
		case *CartesianPoint:
			newZ = p.Z - offset
			bx, by, bz := CalculateFieldFromLoopCartesian(s.Current, loopRadius, p.X, p.Y, newZ)
			Bi += bx
			Bj += by
			Bk += bz
		case *PolarPoint:
			newZ = p.Z - offset
			br, bphi, bz := CalculateFieldFromLoopPolar(s.Current, loopRadius, p.R, newZ)
			Bi += br
			Bj += bphi
			Bk += bz
		default:
			panic(fmt.Sprintf("Unsupported point type: %T", p))
		}
	}
	return
}

func generatePoints(rMin, rMax, phiMin, phiMax, zMin, zMax float64, nr, nphi, nz int) <-chan FieldPoint {
	pointStream := make(chan FieldPoint)

	go func() {
		defer close(pointStream)

		for r := rMin; r <= rMax; r += (rMax - rMin) / float64(nr) {
			for phi := phiMin; phi <= phiMax; phi += (phiMax - phiMin) / float64(nphi) {
				for z := zMin; z <= zMax; z += (zMax - zMin) / float64(nz) {
					pointStream <- NewPolarPoint(r, phi, z)
				}
			}
		}
	}()
	return pointStream
}

func (s *Solenoid) calcField(pointStream <-chan FieldPoint) <-chan FieldPoint {
	resultField := make(chan FieldPoint)

	go func() {
		defer close(resultField)
		for fp := range pointStream {
			zStart := s.CentrePos - s.Length/2
			height := s.Router - s.Rinner

			loopSeparation := s.Length / float64(s.Nturns)
			layerSeparation := height / float64(s.Nlayers)

			Bi, Bj, Bk := s.calculateFieldOverLayersSeq(fp, layerSeparation, loopSeparation, zStart, s.Rinner+0.5*layerSeparation)
			switch p := fp.(type) {
			case *CartesianPoint:
				p.SetFieldCartesian(Bi, Bj, Bk)
			case *PolarPoint:
				p.SetFieldPolar(Bi, Bj, Bk)
			default:
				panic(fmt.Sprintf("Unsupported point type: %T", p))
			}

			resultField <- fp
		}
	}()

	return resultField
}

func accumulateResults(resultChans ...<-chan FieldPoint) <-chan FieldPoint {
	accumulatedResults := make(chan FieldPoint)

	var wg sync.WaitGroup
	wg.Add(len(resultChans))

	for _, c := range resultChans {
		go func(c <-chan FieldPoint) {
			defer wg.Done()
			for fp := range c {
				accumulatedResults <- fp
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(accumulatedResults)
	}()

	return accumulatedResults
}

func (s *Solenoid) CalculateFieldWithWorkers(rMin, rMax, phiMin, phiMax, zMin, zMax float64, nr, nphi, nz, numWorkers int) *Field {
	fieldStream := generatePoints(rMin, rMax, phiMin, phiMax, zMin, zMax, nr, nphi, nz)

	// fan out
	workerChannels := make([]<-chan FieldPoint, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerChannels[i] = s.calcField(fieldStream)
	}

	// fan in
	resultChan := accumulateResults(workerChannels...)

	field := Field{
		Points: make([]FieldPoint, 0, nr*nphi*nz),
	}

	for fp := range resultChan {
		field.Points = append(field.Points, fp)
	}

	return &field
}
