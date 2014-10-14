package main

import (
	"github.com/go-gl/mathgl/mgl32"
	three "github.com/tobscher/go-three"
)

const (
	fov    = 45.0
	width  = 640
	height = 480
	near   = 0.1
	far    = 100
)

func main() {
	scene := three.NewScene()
	camera := three.NewPerspectiveCamera(fov, width/height, near, far)
	camera.Position = mgl32.Vec3{4.0, 3.0, 3.0}
	camera.LookAt(mgl32.Vec3{0.0, 0.0, 0.0})

	// geometry := three.NewTriangleGeometry(
	// 	mgl32.Vec3{-1.0, -1.0, 0.0},
	// 	mgl32.Vec3{1.0, -1.0, 0.0},
	// 	mgl32.Vec3{0.0, 1.0, 0.0},
	// )
	// geometry2 := three.NewTriangleGeometry(
	// 	mgl32.Vec3{-5.0, -5.0, 0.0},
	// 	mgl32.Vec3{-3.0, -5.0, 0.0},
	// 	mgl32.Vec3{0.0, -3.0, 0.0},
	// )

	box := three.NewCubeGeometry(1)
	plane := three.NewBoxGeometry(20, 0.1, 20)

	// red := three.NewMeshBasicMaterial().
	// 	SetColor(three.Color{1.0, 0.0, 0.0})

	green := three.NewMeshBasicMaterial().
		SetColor(three.Color{0.0, 1.0, 0.0})

	blue := three.NewMeshBasicMaterial().
		SetColor(three.Color{0.0, 0.0, 1.0}).
		SetWireframe(false)

	// mesh := three.NewMesh(geometry, red)
	// mesh2 := three.NewMesh(geometry2, green)
	mesh3 := three.NewMesh(box, blue)
	meshPlane := three.NewMesh(plane, green)
	// scene.Add(&mesh)
	// scene.Add(&mesh2)
	scene.Add(&mesh3)
	scene.Add(&meshPlane)

	renderer, err := three.NewRenderer(width, height, "Application Name")
	if err != nil {
		panic(err)
	}

	var i float32 = 0
	for !renderer.ShouldClose() {
		i += 0.01
		// mesh.Rotation.X += 0.01
		// mesh.Rotation.Y += 0.02
		mesh3.SetPosition(0, i, 0)

		renderer.Render(scene, camera)
	}

	renderer.OpenGLSentinel()
}
