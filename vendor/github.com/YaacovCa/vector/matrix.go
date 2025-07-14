package vector

import "math"

type matrix3x3 struct {
	elements [3][3]float64
}

func NewMatrix3x3(elements [3][3]float64) matrix3x3 {
	return matrix3x3{elements: elements}
}

func (m matrix3x3) MultiplyVector(v Vector3) Vector3 {
	return Vector3{
		X: m.elements[0][0]*v.X + m.elements[0][1]*v.Y + m.elements[0][2]*v.Z,
		Y: m.elements[1][0]*v.X + m.elements[1][1]*v.Y + m.elements[1][2]*v.Z,
		Z: m.elements[2][0]*v.X + m.elements[2][1]*v.Y + m.elements[2][2]*v.Z,
	}
}

// Handy matrices for vector operations

func IdentityMatrix3x3() matrix3x3 {
	return NewMatrix3x3([3][3]float64{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	})
}

func RotationMatrix3x3Z(angle float64) matrix3x3 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return NewMatrix3x3([3][3]float64{
		{cos, -sin, 0},
		{sin, cos, 0},
		{0, 0, 1},
	})
}

func RotationMatrix3x3Y(angle float64) matrix3x3 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return NewMatrix3x3([3][3]float64{
		{cos, 0, sin},
		{0, 1, 0},
		{-sin, 0, cos},
	})
}

func RotationMatrix3x3X(angle float64) matrix3x3 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return NewMatrix3x3([3][3]float64{
		{1, 0, 0},
		{0, cos, -sin},
		{0, sin, cos},
	})
}