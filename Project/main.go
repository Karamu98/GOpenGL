package main

import (
	"fmt"
	_ "image/png"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const winWidth = 1280
const winHeight = 720

var objShader *Shader

var sceneLight Light = Light{
	pos:    mgl32.Vec3{5, 5, 5},
	colour: mgl32.Vec3{1, 1, 1},
}

// Light ... A simple light
type Light struct {
	pos    mgl32.Vec3
	colour mgl32.Vec3
}

func main() {

	runtime.LockOSThread()

	fmt.Println("Hello World")
	startUp()
	run()
	shutDown()
}

func startUp() {

	// Try initialise GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialise GLFW:", err)
	}

	// Try create a window
	window, err := glfw.CreateWindow(winWidth, winHeight, "Demo", nil, nil)
	if err != nil {
		panic(err)
	}

	// Focus this window into context
	window.MakeContextCurrent()

	// Try initialise OpenGL
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version", version)

	objShader = createShader("res/shaders/blinnphong.glsl")

	bindShader(objShader)
}

func shutDown() {
	glfw.Terminate()
}

func run() {

	window := glfw.GetCurrentContext()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		if bindShader(objShader) {
			setVec3(objShader, "gLight.pos", sceneLight.pos)
			setVec3(objShader, "gLight.colour", sceneLight.colour)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
