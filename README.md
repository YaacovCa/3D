# 3D

**3D** is a lightweight geometry engine written in Go. It allows you to create and manipulate basic 3D meshes like **cubes** and **rings**, and export them to `.obj` files.

## ✨ Features

- Generate basic 3D meshes:
  - `buildCube()`: creates a cube mesh
  - `buildRing()`: creates a ring mesh
- Apply geometric transformations to meshes:
  - `Move(dx, dy, dz)`
  - `Scale(sx, sy, sz)`
  - `Rotate(rx, ry, rz)`
- Export meshes to `.obj` format (compatible with Blender, Unity, etc.)

## 🧱 Core Structure

The core type is `Mesh`, defined in the `geometry` package.  
It represents a 3D object and supports transformation methods directly:

```go
mesh := BuildCube(1.0)
mesh.Move(0, 2, 0)
mesh.Scale(2, 1, 1)
mesh.Rotate(0, 90, 0)
mesh.SaveAsOBJ("output/cube.obj")
