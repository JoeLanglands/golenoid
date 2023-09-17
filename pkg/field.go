package golenoid

// Field represents a magnetic field as a slice of FieldPoints.
type Field struct {
	Points []FieldPoint
}

// NewField creates a new Field with n FieldPoints.
func NewField(n int) *Field {
	return &Field{
		Points: make([]FieldPoint, n),
	}
}

// TODO @JoeLanglands figure out a better api to create fields from grids.
func NewFieldGridPolar(rMin, rMax, phiMin, phiMax, zMin, zMax float64, nr, nphi, nz int) *Field {
	field := &Field{Points: make([]FieldPoint, 0)}
	for r := rMin; r <= rMax; r += (rMax - rMin) / float64(nr) {
		for phi := phiMin; phi <= phiMax; phi += (phiMax - phiMin) / float64(nphi) {
			for z := zMin; z <= zMax; z += (zMax - zMin) / float64(nz) {
				field.Points = append(field.Points, NewPolarPoint(r, phi, z))
			}
		}
	}

	return field
}

func (f *Field) WriteToFile() error {
	return nil
}
