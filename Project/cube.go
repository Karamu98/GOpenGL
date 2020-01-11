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

	// Right face
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
	vbo    uint32
	vao    uint32
	objMat mgl32.Mat4
}

func initCube() *Cube {
	newCube := &Cube{}

	gl.GenVertexArrays(1, &newCube.vao)
	gl.GenBuffers(1, &newCube.vbo)

	gl.BindVertexArray(newCube.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, newCube.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexData)*4, gl.Ptr(vertexData), gl.STATIC_DRAW)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 8*4, nil)

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 8*4, gl.PtrOffset(3*4))

	gl.EnableVertexAttribArray(2)
	gl.VertexAttribPointer(2, 2, gl.FLOAT, false, 8*4, gl.PtrOffset(6*4))

	newCube.objMat = mgl32.Ident4()

	return newCube
}

func drawCube(cubeToDraw *Cube) {
	gl.BindVertexArray(cubeToDraw.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 36)
	gl.BindVertexArray(0)
}
