package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func keyPress(bodies []*Body) {
	if rl.IsKeyPressed(rl.KeyG) {
		for _, body := range bodies {
			body.Randomize()
		}
	}
}

func main() {
	rl.InitWindow(800, 450, "Matrix Rotation - Ben Schreiber")

	rl.SetTargetFPS(60)

	rl.DisableCursor()

	camera := NewFirstPersonCamera()

	bodies := []*Body{
		NewBody(rl.NewVector3(0.0, 2.0, 0.0)),
	}

	for !rl.WindowShouldClose() {

		camera.Update()
		for _, body := range bodies {
			body.Rotate(body.AngularVelocity)
		}

		keyPress(bodies)

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera.Camera)

		DrawScene()
		for _, body := range bodies {
			body.Draw()
		}

		rl.EndMode3D()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
