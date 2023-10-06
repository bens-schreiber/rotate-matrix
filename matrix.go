package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Body struct {
	Position        rl.Vector3
	Angle           rl.Vector3
	Size            rl.Vector3
	AngularVelocity rl.Vector3
	Color           rl.Color
}

func NewBody(position rl.Vector3) *Body {
	return &Body{Position: position, Angle: rl.Vector3Zero(), Size: rl.Vector3One(), AngularVelocity: rl.NewVector3(1, 1, 1), Color: rl.Blue}
}

func (body *Body) Draw() {

	// Enter a matrix context to perform matrix operations
	rl.PushMatrix()

	// Put the matrix in the body's position
	rl.Translatef(body.Position.X, body.Position.Y, body.Position.Z)

	// Rotate the matrix
	rl.Rotatef(body.Angle.X, 1.0, 0.0, 0.0)
	rl.Rotatef(body.Angle.Y, 0.0, 1.0, 0.0)
	rl.Rotatef(body.Angle.Z, 0.0, 0.0, 1.0)

	rl.DrawCube(rl.Vector3Zero(), body.Size.X, body.Size.Y, body.Size.Z, body.Color)
	rl.DrawCubeWires(rl.Vector3Zero(), body.Size.X, body.Size.Y, body.Size.Z, rl.Black)

	// Leave the matrix context
	rl.PopMatrix()

	drawAngularVelocityIndicators(body)
}

func drawAngularVelocityIndicators(body *Body) {
	const indicatorLength = 2
	rl.PushMatrix()

	rl.Translatef(body.Position.X, body.Position.Y, body.Position.Z)
	rl.Rotatef(body.Angle.X, 1.0, 0.0, 0.0)
	rl.Rotatef(body.Angle.Y, 0.0, 1.0, 0.0)
	rl.Rotatef(body.Angle.Z, 0.0, 0.0, 1.0)

	// X indicator
	rl.DrawCube(rl.NewVector3(-0.5*(indicatorLength*body.AngularVelocity.X), 0, 0), indicatorLength*body.AngularVelocity.X, 0.1, 0.1, rl.Red)

	// Y indicator
	rl.DrawCube(rl.NewVector3(0, -0.5*(indicatorLength*body.AngularVelocity.Y), 0), 0.1, indicatorLength*body.AngularVelocity.Y, 0.1, rl.Green)

	// Z indicator
	rl.DrawCube(rl.NewVector3(0, 0, -0.5*(indicatorLength*body.AngularVelocity.Z)), 0.1, 0.1, indicatorLength*body.AngularVelocity.Z, rl.Blue)

	rl.PopMatrix()
}

func (body *Body) Rotate(angle rl.Vector3) {
	body.Angle = rl.Vector3Add(body.Angle, angle)
}

func (body *Body) Translate(position rl.Vector3) {
	body.Position = position
}

func (body *Body) Resize(size rl.Vector3) {
	body.Size = size
}

func (body *Body) Velocity(velocity rl.Vector3) {
	body.AngularVelocity = velocity
}

func (body *Body) Randomize() {
	body.Translate(rl.NewVector3(float32(rl.GetRandomValue(0, 16)), float32(rl.GetRandomValue(3, 5)), float32(rl.GetRandomValue(0, 16))))
	body.Resize(rl.NewVector3(float32(rl.GetRandomValue(1, 4)), float32(rl.GetRandomValue(1, 4)), float32(rl.GetRandomValue(1, 4))))
	body.Velocity(rl.NewVector3(float32(rl.GetRandomValue(-3, 3)), float32(rl.GetRandomValue(-3, 3)), float32(rl.GetRandomValue(-3, 3))))
	body.Color = rl.NewColor(uint8(rl.GetRandomValue(0, 255)), uint8(rl.GetRandomValue(0, 255)), uint8(rl.GetRandomValue(0, 255)), 255)
}
