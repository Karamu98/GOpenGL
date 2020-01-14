package main

import "github.com/go-gl/mathgl/mgl32"

type light struct {
	position mgl32.Vec3
	colour   mgl32.Vec3
}

func (light *light) Draw(shader *Shader) {
	shader.SetVec3("gLight.pos", light.position)
	shader.SetVec3("gLight.colour", light.colour)
}

// PointLight ... A point light for rendering
type PointLight struct {
	light       light
	attenuation float32
}

// Draw ... Sets uniforms with light data
func (light *PointLight) Draw(shader *Shader) {
	light.light.Draw(shader)
	shader.SetFloat("gLight.attenuation", light.attenuation)
}
