package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawScene() {
	rl.DrawPlane(rl.Vector3{X: 0.0, Y: 0.0, Z: 0.0}, rl.NewVector2(32.0, 32.0), rl.LightGray)
}
