package abs

import (
	"math"
)

type Vertex struct {
	X float64
	Y float64
}

func NewVertex() *Vertex {
	return &Vertex{}
}

func SetVertexValue(v *Vertex, x, y float64) Vertex {
	v.X = x
	v.Y = y

	return *v
}

// same as Abs(v Vertex) float64
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
