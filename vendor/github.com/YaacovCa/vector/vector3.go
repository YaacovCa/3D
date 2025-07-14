package vector

import "math"

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// Basic vector operations for 3D vectors

func (v Vector3) Add(o Vector3) Vector3 {
	return Vector3{
		X: v.X + o.X,
		Y: v.Y + o.Y,
		Z: v.Z + o.Z,
	}
}

func (v Vector3) Sub(o Vector3) Vector3 {
	return Vector3{
		X: v.X - o.X,
		Y: v.Y - o.Y,
		Z: v.Z - o.Z,
	}
}

func (v Vector3) Scale(s float64) Vector3 {
	return Vector3{
		X: v.X * s,
		Y: v.Y * s,
		Z: v.Z * s,
	}
}

func (v Vector3) Dot(o Vector3) float64 {
	return v.X * o.X + v.Y * o.Y + v.Z * o.Z
}

func (v Vector3) Magnitude() float64 {
	return (math.Sqrt(v.X*v.X + v.Y * v.Y + v.Z * v.Z))
}

func (v Vector3) Normalize() Vector3 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector3{X: 0, Y: 0, Z: 0}
	}
	return Vector3{
		X: v.X / mag,
		Y: v.Y / mag,
		Z: v.Z / mag,
	}
}

func (v Vector3) Cross(o Vector3) Vector3 {
	return Vector3{
		X: v.Y*o.Z - v.Z*o.Y,
		Y: v.Z*o.X - v.X*o.Z,
		Z: v.X*o.Y - v.Y*o.X,
	}
}

func (v Vector3) AngleBetween(o Vector3) float64 {
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

func (v Vector3) MultipliedByMatrix(m matrix3x3) Vector3 {
	return m.MultiplyVector(v)
}

// Handy type for 3D vectors

func Vector3Zero() Vector3 {
	return Vector3{0, 0, 0}
}

func Vector3One() Vector3 {
	return Vector3{1, 1, 1}
}

func Vector3Up() Vector3 {
	return Vector3{0, 1, 0}
}

func Vector3Down() Vector3 {
	return Vector3{0, -1, 0}
}

func Vector3Right() Vector3 {
	return Vector3{1, 0, 0}
}

func Vector3Left() Vector3 {
	return Vector3{-1, 0, 0}
}

func Vector3Forward() Vector3 {
	return Vector3{0, 0, 1}
}

func Vector3Back() Vector3 {
	return Vector3{0, 0, -1}
}
