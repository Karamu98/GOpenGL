package main

import (
	"fmt"
	_ "image/png"
	"log"
	"runtime"

	"./shader"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const winWidth = 1280
const winHeight = 720

var objShader *shader.Shader
var cube *Cube
var cam *Camera

var sceneLight Light = Light{
	pos:    mgl32.Vec3{5, 5, 5},
	colour: mgl32.Vec3{255, 255, 255},
}

// Light ... A simple light
type Light struct {
	pos    mgl32.Vec3
	colour mgl32.Vec3
}

// Camera .. A simple camera
type Camera struct {
	obj        mgl32.Mat4
	projection mgl32.Mat4
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

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(0.2, 0.2, 0.2, 1.0)

	objShader = shader.Create("res/shaders/blinnphong.glsl")

	objShader.Bind()

	cam = &Camera{}

	cam.obj = mgl32.LookAtV(mgl32.Vec3{5, 5, 5}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
	cam.projection = mgl32.Perspective(mgl32.DegToRad(60.0), float32(winWidth)/winHeight, 0.1, 100.0)

	projView := cam.projection.Mul4(cam.obj)

	camPos := mgl32.Vec3{
		cam.obj.Row(3).X(),
		cam.obj.Row(3).Y(),
		cam.obj.Row(3).Z(),
	}

	objShader.SetMat4("camProjView", projView)
	objShader.SetVec3("gCamPos", camPos)

	cube = initCube()
}

func shutDown() {
	glfw.Terminate()
}

func run() {

	window := glfw.GetCurrentContext()

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		if objShader.Bind() {
			objShader.SetVec3("gLight.pos", sceneLight.pos)
			objShader.SetVec3("gLight.colour", sceneLight.colour)

			objShader.SetFloat("gGamma", 1.8)

			objShader.SetMat4("objMatrix", cube.objMat)
			cube.draw()
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
