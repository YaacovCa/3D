package vector

import "math"

type Vector2 struct {
	X float64
	Y float64
}

// Basic vector operations for 2D vectors

func (v Vector2) Add(o Vector2) Vector2 {
	return Vector2{
		X: v.X + o.X,
		Y: v.Y + o.Y,
	}
}

func (v Vector2) Sub(o Vector2) Vector2 {
	return Vector2{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (v Vector2) Scale(s float64) Vector2 {
	return Vector2{
		X: v.X * s,
		Y: v.Y * s,
	}
}

func (v Vector2) Dot(o Vector2) float64 {
	return v.X * o.X + v.Y * o.Y
}

func (v Vector2) Magnitude() float64 {
	return (math.Sqrt(v.X*v.X + v.Y * v.Y))
}

func (v Vector2) Normalize() Vector2 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector2{X: 0, Y: 0}
	}
	return Vector2{
		X: v.X / mag,
		Y: v.Y / mag,
	}
}

func (v Vector2) AngleBetween(o Vector2) float64 {
	dot := v.Dot(o)
	magV := v.Magnitude()
	magO := o.Magnitude()
	if magV == 0 || magO == 0 {
		return 0
	}
	cosTheta := dot / magO * magV
	// Clamp cosTheta to the range [-1, 1] to avoid NaN from
	cosTheta = math.Max(-1.0, math.Min(1.0, cosTheta))
	return math.Acos(cosTheta)
}

// Handy type for 2D vectors

func Vector2Zero() Vector2 {
	return Vector2{0, 0}
}

func Vector2One() Vector2 {
	return Vector2{1, 1}
}

func Vector2Up() Vector2 {
	return Vector2{0, 1}
}

func Vector2Down() Vector2 {
	return Vector2{0, -1}
}

func Vector2Right() Vector2 {
	return Vector2{1, 0}
}

func Vector2Left() Vector2 {
	return Vector2{-1, 0}
}
