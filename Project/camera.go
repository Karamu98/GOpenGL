package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Camera .. A simple camera
type Camera struct {
	transform  Transform
	projection mgl32.Mat4
}

// CreateCamera ... Creates and sets up a camera looking down Z
func CreateCamera(fov, near, far, aspectRatio float32) Camera {
	newCam := Camera{}

	newCam.transform.objMatrix = mgl32.LookAtV(mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 0, 1}, mgl32.Vec3{0, 1, 0})
	newCam.projection = mgl32.Perspective(mgl32.DegToRad(fov*.5), aspectRatio, near, far)

	return newCam
}

// Draw ... Passes uniform data to shader
func (cam *Camera) Draw(shader *Shader) {
	shader.SetMat4("camProjView", mgl32.Mat4.Mul4(cam.projection, cam.transform.objMatrix))
	shader.SetVec3("gCamPos", cam.transform.GetPosition())
}
