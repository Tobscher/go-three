package main

import (
	"log"

	three "github.com/tobscher/go-three"
	"github.com/tobscher/go-three/geometries"
)

const (
	fov    = 45.0
	width  = 640
	height = 480
	near   = 0.1
	far    = 100
)

func main() {
	window, err := three.NewWindow(width, height, "Example - Wireframe Cube")
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := three.NewRenderer(window)
	if err != nil {
		log.Fatal(err)
	}

	scene := three.NewScene()
	camera := three.NewPerspectiveCamera(fov, width/height, near, far)
	camera.Transform.SetPosition(4.0, 3.0, 4.0)
	camera.Transform.LookAt(0, 0, 0)

	box := geometries.NewCube(1)
	blue := three.NewBasicMaterial()
	blue.SetColor(&three.Color{0.0, 0.0, 1.0})
	blue.SetWireframe(true)

	mesh := three.NewMesh(box, blue)

	scene.Add(&mesh)

	blue.SetColor(&three.Color{1.0, 0.0, 0.0})

	for !window.ShouldClose() {
		mesh.Transform.RotateX(0.01)
		mesh.Transform.RotateY(0.02)

		renderer.Render(scene, camera)
	}

	renderer.Unload(scene)

	renderer.OpenGLSentinel()
}
