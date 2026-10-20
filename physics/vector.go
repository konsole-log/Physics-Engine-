package physics

import (
	"fmt"
	"math"
)

type Vector3 struct {
	X, Y, Z float64
}

func (v *Vector3) Invert() {
	v.X = -v.X
	v.Y = -v.Y
	v.Z = -v.Z
}

func (v *Vector3) Normalize() {
	l := v.Magnitude()
	if l > 0 {
		*v = v.Scale(1 / l)
	}
}

func (v Vector3) Dot(w Vector3) float64 {
	return (v.X*w.X + v.Y*w.Y + v.Z*w.Z)
}

func (v Vector3) Cross(w Vector3) Vector3 {
	return Vector3{v.Y*w.Z - v.Z*w.Y, v.Z*w.X - v.X*w.Z, v.X*w.Y - v.Y*w.X}
}

func (v Vector3) Add(w Vector3) Vector3 {
	return Vector3{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func (v Vector3) AddScaledVector(w Vector3, scale float64) Vector3 {
	return v.Add(w.Scale(scale))
}

func (v Vector3) Subtract(w Vector3) Vector3 {
	return Vector3{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v Vector3) SubtractScaledVector(w Vector3, scale float64) Vector3 {
	return v.Add(w.Scale(scale))
}

func (v Vector3) Magnitude() float64 {
	return math.Sqrt(v.SquareMagnitude())
}

func (v Vector3) SquareMagnitude() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector3) Scale(value float64) Vector3 {
	return Vector3{v.X * value, v.Y * value, v.Z * value}
}

func (v Vector3) Print() {
	fmt.Printf("[%.3f, %.3f, %.3f]\n", v.X, v.Y, v.Z)
}

func NewVector3(x, y, z float64) Vector3 {
	return Vector3{X: x, Y: y, Z: z}
}
