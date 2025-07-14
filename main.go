package main

import (
	"math"
	"os"

	"github.com/YaacovCa/vector"

	"github.com/YaacovCa/3d/geometry"
)

func buildRing(sides int, radius float64, height float64, thickness float64, doubleSided bool) geometry.Mesh {

	tris := make([]geometry.Tri, 0)
	vertices := make([]vector.Vector3, 0)

	angleIncrement := (2.0 * math.Pi) / float64(sides)

	for i := 0; i < sides; i++ {
		v := vector.Vector3{
			X: radius * math.Cos(float64(i)*angleIncrement),
			Y: 0,
			Z: radius * math.Sin(float64(i)*angleIncrement),
		}
		vertices = append(vertices, v, v.Add(vector.Vector3{X: 0, Y: height, Z: 0}))
	}

	// add thickness to the ring
	for _, v := range vertices {
		vertices = append(vertices, vector.Vector3{X: v.X * (1.0 + thickness), Y: v.Y, Z: v.Z * (1.0 + thickness)})
	}

	for i := 1; i < sides; i++ {
		curVert := i * 2
		tris = append(tris, geometry.Tri{P1: curVert - 1, P2: curVert, P3: curVert + 1})
		tris = append(tris, geometry.Tri{P1: curVert - 2, P2: curVert, P3: curVert - 1})
	}

	// Connect the last vertex to the first
	tris = append(tris, geometry.Tri{P1: 1, P2: sides*2 - 1, P3: 0})
	tris = append(tris, geometry.Tri{P1: 0, P2: sides*2 - 1, P3: sides*2 - 2})

	for i := sides + 1; i < sides*2; i++ {
		curVert := i * 2
		tris = append(tris, geometry.Tri{P1: curVert + 1, P2: curVert, P3: curVert - 1})
		tris = append(tris, geometry.Tri{P1: curVert - 1, P2: curVert, P3: curVert - 2})
	}
	// Connect the last vertex to the first
	tris = append(tris, geometry.Tri{P1: 1 + sides*2, P2: 0 + sides*2, P3: sides*4 - 1})
	tris = append(tris, geometry.Tri{P1: 0 + sides*2, P2: sides*4 - 2, P3: sides*4 - 1})

	for i := 1; i < sides; i++ {
		curVert := i * 2
		tris = append(tris, geometry.Tri{P1: curVert, P2: curVert - 2, P3: curVert + sides*2})
		tris = append(tris, geometry.Tri{P1: curVert - 2, P2: curVert + sides*2 - 2, P3: curVert + sides*2})
		tris = append(tris, geometry.Tri{P1: curVert + 1, P2: curVert + sides*2 + 1, P3: curVert -1})
		tris = append(tris, geometry.Tri{P1: curVert - 1, P2: curVert + sides*2 + 1, P3: curVert + sides*2 - 1})
	}

	tris = append(tris, geometry.Tri{P1: 0, P2: sides*2 - 2, P3: sides*4 - 2})
	tris = append(tris, geometry.Tri{P1: sides*4 - 2, P2: sides*2, P3: 0})
	tris = append(tris, geometry.Tri{P1: 1, P2: sides*4 - 1, P3: sides*2 - 1})
	tris = append(tris, geometry.Tri{P1: sides*4 - 1, P2: 1, P3: sides*2 + 1})

	if doubleSided {
		for _, tri := range tris {
			tris = append(tris, geometry.Tri{
				P1: tri.P3, P2: tri.P2, P3: tri.P1,
			})
		}
	}

	return geometry.Mesh{Triangles: tris, Vertices: vertices}
}

func buildCube(width float64, heigt float64, depth float64, doubleSided bool) geometry.Mesh {
	tris := make([]geometry.Tri, 0)
	vertices := make([]vector.Vector3, 0)
	// Bottom vertices
	vertices = append(vertices, vector.Vector3{
		X: -width / 2, Y: 0, Z: -depth / 2,
	})
	vertices = append(vertices, vector.Vector3{
		X: +width / 2, Y: 0, Z: -depth / 2,
	})
	vertices = append(vertices, vector.Vector3{
		X: +width / 2, Y: 0, Z: +depth / 2,
	})
	vertices = append(vertices, vector.Vector3{
		X: -width / 2, Y: 0, Z: +depth / 2,
	})

	// Top vertices
	vertices = append(vertices, vector.Vector3{
		X: -width / 2, Y: +heigt, Z: -depth / 2,
	})
	vertices = append(vertices, vector.Vector3{
		X: +width / 2, Y: +heigt, Z: -depth / 2,
	})
	vertices = append(vertices, vector.Vector3{
		X: +width / 2, Y: +heigt, Z: +depth / 2,
	})
	vertices = append(vertices, vector.Vector3{
		X: -width / 2, Y: +heigt, Z: +depth / 2,
	})

	// Bottom face
	tris = append(tris, geometry.Tri{P1: 0, P2: 1, P3: 2})
	tris = append(tris, geometry.Tri{P1: 2, P2: 3, P3: 0})

	// Left face
	tris = append(tris, geometry.Tri{P1: 0, P2: 7, P3: 4})
	tris = append(tris, geometry.Tri{P1: 7, P2: 0, P3: 3})

	// Front face
	tris = append(tris, geometry.Tri{P1: 2, P2: 7, P3: 3})
	tris = append(tris, geometry.Tri{P1: 7, P2: 2, P3: 6})

	// Right face
	tris = append(tris, geometry.Tri{P1: 2, P2: 1, P3: 5})
	tris = append(tris, geometry.Tri{P1: 5, P2: 6, P3: 2})

	// Back face
	tris = append(tris, geometry.Tri{P1: 1, P2: 0, P3: 4})
	tris = append(tris, geometry.Tri{P1: 4, P2: 5, P3: 1})

	// Top face
	tris = append(tris, geometry.Tri{P1: 4, P2: 6, P3: 5})
	tris = append(tris, geometry.Tri{P1: 6, P2: 4, P3: 7})

	if doubleSided {
		for _, tri := range tris {
			tris = append(tris, geometry.Tri{
				P1: tri.P3, P2: tri.P2, P3: tri.P1,
			})
		}
	}

	return geometry.Mesh{Triangles: tris, Vertices: vertices}
}

func main() {
	/*cube := buildCube(7, 10, 6, true)
	cube.ToOBJ(os.Stdout)*/

	ring := buildRing(50, 3, 8, 0.1, true)
	ring = ring.Rotate("X", 45)
	ring = ring.Move(1, 2, 3)
	ring = ring.Scale(0.5)
	ring.ToOBJ(os.Stdout)

}
