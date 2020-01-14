package main

import "github.com/go-gl/mathgl/mgl32"

// Transform ... Defines an object in space
type Transform struct {
	objMatrix mgl32.Mat4
}

// SetPosition ... Sets the position of the object
func (transform *Transform) SetPosition(newPos mgl32.Vec3) {
	transform.objMatrix.SetRow(3, mgl32.Vec4{newPos.X(), newPos.Y(), newPos.Z()})
}

// GetPosition ... Gets the position of the object
func (transform *Transform) GetPosition() mgl32.Vec3 {
	pos := transform.objMatrix.Row(3)
	return mgl32.Vec3{pos.X(), pos.Y(), pos.Z()}
}

// LookAt ... Makes the object face target
func (transform *Transform) LookAt(target mgl32.Vec3) {
	transform.objMatrix = mgl32.LookAtV(transform.GetPosition(), target, mgl32.Vec3{0, 1, 0})
}

// Rotate ... Rotates the object about axis by angle
func (transform *Transform) Rotate(angle float32, axis mgl32.Vec3) {
	transform.objMatrix = mgl32.HomogRotate3D(angle, axis)
}
