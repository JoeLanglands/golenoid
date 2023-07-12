package golenoid

// Field represents a magnetic field as a slice of FieldPoints.
type Field struct {
	Points []FieldPoint
}

func NewField(n int) *Field {
	return &Field{
		Points: make([]FieldPoint, n),
	}
}
