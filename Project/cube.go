package main

import (
	"github.com/go-gl/mathgl/mgl32"
)

vertexData := [...]float32{
	// Positions			// Normals				// UV's

	// Back face
	 1.0f, -1.0, -1.0,	0.0,  0.0, -1.0,		0.0,  0.0,
	-1.0, -1.0, -1.0,	0.0,  0.0, -1.0,		1.0,  0.0,
	-1.0,  1.0, -1.0,	0.0,  0.0, -1.0,		1.0,  1.0,
	-1.0,  1.0, -1.0,	0.0,  0.0, -1.0,		1.0,  1.0,
	 1.0,  1.0, -1.0,	0.0,  0.0, -1.0,		0.0,  1.0,
	 1.0, -1.0, -1.0,	0.0,  0.0, -1.0,		0.0,  0.0,
										 
	// Front face						 
	-1.0, -1.0,  1.0,	0.0,  0.0,  1.0,		0.0,  0.0,
	 1.0, -1.0,  1.0,	0.0,  0.0,  1.0,		1.0,  0.0,
	 1.0,  1.0,  1.0,	0.0,  0.0,  1.0,		1.0,  1.0,
	 1.0,  1.0,  1.0,	0.0,  0.0,  1.0,		1.0,  1.0,
	-1.0,  1.0,  1.0,	0.0,  0.0,  1.0,		0.0,  1.0,
	-1.0, -1.0,  1.0,	0.0,  0.0,  1.0,		0.0,  0.0,
										 
	// Left face						 
	-1.0,  1.0,  1.0,	-1.0,  0.0,  0.0,		1.0,  1.0,
	-1.0,  1.0, -1.0,	-1.0,  0.0,  0.0,		0.0,  1.0,
	-1.0, -1.0, -1.0,	-1.0,  0.0,  0.0,		0.0,  0.0,
	-1.0, -1.0, -1.0,	-1.0,  0.0,  0.0,		0.0,  0.0,
	-1.0, -1.0,  1.0,	-1.0,  0.0,  0.0,		1.0,  0.0,
	-1.0,  1.0,  1.0,	-1.0,  0.0,  0.0,		1.0,  1.0,
										 
	// Right face		 
	 1.0, -1.0,  1.0,	1.0,  0.0,  0.0,		0.0,  0.0,
	 1.0, -1.0, -1.0,	1.0,  0.0,  0.0,		1.0,  0.0,
	 1.0,  1.0, -1.0,	1.0,  0.0,  0.0,		1.0,  1.0,
	 1.0,  1.0, -1.0,	1.0,  0.0,  0.0,		1.0,  1.0,
	 1.0,  1.0,  1.0,	1.0,  0.0,  0.0,		0.0,  1.0,
	 1.0, -1.0,  1.0,	1.0,  0.0,  0.0,		0.0,  0.0,
	
	// Bottom face
	-1.0, -1.0, -1.0,	0.0, -1.0,  0.0,		0.0,  1.0,
	 1.0, -1.0, -1.0,	0.0, -1.0,  0.0,		1.0,  1.0,
	 1.0, -1.0,  1.0,	0.0, -1.0,  0.0,		1.0,  0.0,
	 1.0, -1.0,  1.0,	0.0, -1.0,  0.0,		1.0,  0.0,
	-1.0, -1.0,  1.0,	0.0, -1.0,  0.0,		0.0,  0.0,
	-1.0, -1.0, -1.0,	0.0, -1.0,  0.0,		0.0,  1.0,
	
	// Top face
	-1.0,  1.0,  1.0,	0.0,  1.0,  0.0,		0.0,  0.0, 
	 1.0,  1.0,  1.0,	0.0,  1.0,  0.0,		1.0,  0.0, 
	 1.0,  1.0, -1.0,	0.0,  1.0,  0.0,		1.0,  1.0, 
	 1.0,  1.0, -1.0,	0.0,  1.0,  0.0,		1.0,  1.0, 
	-1.0,  1.0, -1.0,	0.0,  1.0,  0.0,		0.0,  1.0, 
	-1.0,  1.0,  1.0,	0.0,  1.0,  0.0,		0.0,  0.0	 
};


type Cube struct {
	vbo uint32
	vao uint32
	objMat mgl32.Mat4
}

func initCube() Cube {
	var newCube Cube 

	gl.GenVertexArrays(1, &newCube.vao)
	gl.GenBuffers(1, &newCube.vbo)

	gl.BindVertexArray(newCube.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, m_vbo);
	gl.BufferData(gl.ARRAY_BUFFER, sizeof(vertexData), vertexData, gl.STATIC_DRAW);

	gl.EnableVertexAttribArray(0);
	gl.VertexAttribPointer(0, 3, gl.FLOAT, gl.FALSE, 8 * sizeof(float), nil);

	gl.EnableVertexAttribArray(1);
	gl.VertexAttribPointer(1, 3, gl.FLOAT, gl.FALSE, 8 * sizeof(float), (void*)(3 * sizeof(float)));

	gl.EnableVertexAttribArray(2);
	gl.VertexAttribPointer(2, 2, gl.FLOAT, gl.FALSE, 8 * sizeof(float), (void*)(6 * sizeof(float)));
}