package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var vertexData = []float32{
	// Positions			// Normals				// UV's

	// Back face
	1.0, -1.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,
	-1.0, -1.0, -1.0, 0.0, 0.0, -1.0, 1.0, 0.0,
	-1.0, 1.0, -1.0, 0.0, 0.0, -1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0, 0.0, 0.0, -1.0, 1.0, 1.0,
	1.0, 1.0, -1.0, 0.0, 0.0, -1.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 0.0, 0.0, -1.0, 0.0, 0.0,

	// Front face
	-1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0,
	1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 1.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 0.0, 1.0, 1.0, 1.0,
	1.0, 1.0, 1.0, 0.0, 0.0, 1.0, 1.0, 1.0,
	-1.0, 1.0, 1.0, 0.0, 0.0, 1.0, 0.0, 1.0,
	-1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 0.0, 0.0,

	// Left face
	-1.0, 1.0, 1.0, -1.0, 0.0, 0.0, 1.0, 1.0,
	-1.0, 1.0, -1.0, -1.0, 0.0, 0.0, 0.0, 1.0,
	-1.0, -1.0, -1.0, -1.0, 0.0, 0.0, 0.0, 0.0,
	-1.0, -1.0, -1.0, -1.0, 0.0, 0.0, 0.0, 0.0,
	-1.0, -1.0, 1.0, -1.0, 0.0, 0.0, 1.0, 0.0,
	-1.0, 1.0, 1.0, -1.0, 0.0, 0.0, 1.0, 1.0,

	// Rightf ace
	1.0, -1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0,
	1.0, -1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 1.0,
	1.0, 1.0, -1.0, 1.0, 0.0, 0.0, 1.0, 1.0,
	1.0, 1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 1.0,
	1.0, -1.0, 1.0, 1.0, 0.0, 0.0, 0.0, 0.0,

	// Bottom face
	-1.0, -1.0, -1.0, 0.0, -1.0, 0.0, 0.0, 1.0,
	1.0, -1.0, -1.0, 0.0, -1.0, 0.0, 1.0, 1.0,
	1.0, -1.0, 1.0, 0.0, -1.0, 0.0, 1.0, 0.0,
	1.0, -1.0, 1.0, 0.0, -1.0, 0.0, 1.0, 0.0,
	-1.0, -1.0, 1.0, 0.0, -1.0, 0.0, 0.0, 0.0,
	-1.0, -1.0, -1.0, 0.0, -1.0, 0.0, 0.0, 1.0,

	// Top face
	-1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 0.0, 0.0,
	1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 1.0, 0.0,
	1.0, 1.0, -1.0, 0.0, 1.0, 0.0, 1.0, 1.0,
	1.0, 1.0, -1.0, 0.0, 1.0, 0.0, 1.0, 1.0,
	-1.0, 1.0, -1.0, 0.0, 1.0, 0.0, 0.0, 1.0,
	-1.0, 1.0, 1.0, 0.0, 1.0, 0.0, 0.0, 0.0,
}

// Cube ... Simple Cube
type Cube struct {
	vbo       uint32
	vao       uint32
	transform Transform
}

// CreateCube ... Loads vertex data into OpenGL and creates cube
func CreateCube() Cube {
	newCube := Cube{}

	gl.GenVertexArrays(1, &newCube.vao)
	gl.GenBuffers(1, &newCube.vbo)

	gl.BindVertexArray(newCube.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, newCube.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData)*4, gl.Ptr(vertexData), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(0))

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))

	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))

	newCube.transform.objMatrix = mgl32.Ident4()

	return newCube
}

func (cube *Cube) draw(shader *Shader) {
	shader.SetMat4("objMatrix", cube.transform.objMatrix)

	gl.BindVertexArray(cube.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
	gl.BindVertexArray(0)
}

// Destroy ... Removes vertex data from OpenGL
func (cube *Cube) Destroy() {
	buffers := []*uint32{&cube.vao, &cube.vbo}
	gl.DeleteBuffers(2, buffers[0])
}
