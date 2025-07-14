package geometry

import (
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/YaacovCa/vector"
)

type Mesh struct {
	Triangles []Tri
	Vertices  []vector.Vector3
}

func (m Mesh) ToOBJ(w io.Writer) error {

	for _, v := range m.Vertices {
		io.WriteString(w, fmt.Sprintf("v %.2f %.2f %.2f\n", v.X, v.Y, v.Z))
	}

	for _, t := range m.Triangles {
		io.WriteString(w, fmt.Sprintf("f %.d %.d %.d\n", t.P1+1, t.P2+1, t.P3+1))
	}
	return nil
}

func (m Mesh) Move(offsetX float64, offsetY float64, offsetZ float64) Mesh {
	for i := range m.Vertices {
		m.Vertices[i] = m.Vertices[i].Add(vector.Vector3{
			X: offsetX, Y: offsetY, Z: offsetZ,
		})
	}
	return Mesh{
		Triangles: m.Triangles,
		Vertices:  m.Vertices,
	}
}

func (m Mesh) Scale(sccale float64) Mesh {
	for i := range m.Vertices {
		m.Vertices[i] = m.Vertices[i].Scale(sccale)
	}
	return Mesh{
		Triangles: m.Triangles,
		Vertices:  m.Vertices,
	}
}

func (m Mesh) Rotate(axis string, angle float64) Mesh {
	axis = strings.ToUpper(axis)
	angle = angle * (math.Pi / 180.0) // Convert degrees to radians

	if axis == "X" {
		for i := range m.Vertices {
			m.Vertices[i] = m.Vertices[i].MultipliedByMatrix(vector.RotationMatrix3x3X(angle))
		}
	}

	if axis == "Y" {
		for i := range m.Vertices {
			m.Vertices[i] = m.Vertices[i].MultipliedByMatrix(vector.RotationMatrix3x3Y(angle))
		}
	}

	if axis == "Z" {
		for i := range m.Vertices {
			m.Vertices[i] = m.Vertices[i].MultipliedByMatrix(vector.RotationMatrix3x3Z(angle))
		}
	}

	return Mesh{
		Triangles: m.Triangles,
		Vertices:  m.Vertices,
	}
}