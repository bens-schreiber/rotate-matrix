package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type FirstPersonCamera struct {
	Camera rl.Camera3D
}

func NewFirstPersonCamera() *FirstPersonCamera {
	var camera rl.Camera3D = rl.Camera3D{}
	camera.Position = rl.Vector3{X: 0.0, Y: 2.0, Z: 4.0}
	camera.Target = rl.Vector3{X: 0.0, Y: 2.0, Z: 0.0}
	camera.Up = rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0}
	camera.Fovy = 60.0
	camera.Projection = rl.CameraPerspective
	return &FirstPersonCamera{Camera: camera}
}

func (camera *FirstPersonCamera) Update() {
	rl.UpdateCamera(&camera.Camera, rl.CameraFirstPerson)
}
