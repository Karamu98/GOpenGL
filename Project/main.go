package main

import (
	_ "image/png"
	"runtime"

	util "./Utilities"
	logger "./Utilities/Logger"
	"./shader"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const winWidth = 1280
const winHeight = 720

var objShader *shader.Shader
var cube Cube
var cam Camera
var texture *Texture

// FBO
var defaultFBO uint32
var defaultColourTex uint32
var depthRender uint32

var sceneLight PointLight = PointLight{
	light: light{
		position: mgl32.Vec3{5, 5, -5},
		colour:   mgl32.Vec3{1, 1, 1},
	},
	attenuation: 0.0,
}

func main() {
	runtime.LockOSThread()

	if startUp() == false {
		shutDown()
		return
	}
	run()
	shutDown()
}

func startUp() bool {
	// OpenGL Initialisation
	{
		// Try initialise GLFW
		if err := glfw.Init(); err != nil {
			panic(err)
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

		// Print version
		version := gl.GoStr(gl.GetString(gl.VERSION))
		logger.Infof("OpenGL Version: %v\n", version)

		// Globals
		gl.Enable(gl.DEPTH_TEST)
		gl.DepthFunc(gl.LESS)
		gl.Enable(gl.CULL_FACE)
		gl.CullFace(gl.BACK)
		gl.ClearColor(0.2, 0.2, 0.2, 1.0)
	}

	// Create and bind shader
	objShader = shader.Create("res/shaders/simple.glsl")
	if !objShader.Bind() {
		logger.Errorln("Failed to create shader.")
		return false
	}

	// Create and bind texture
	texture = CreateTexture("res/textures/test.png")

	objShader.SetFloat("gGamma", 1.8)
	objShader.SetInt("gMaterial.texture", 0)
	objShader.SetFloat("gMaterial.spec", 32.0)

	cam = createCamera(100, 0.1, 100)
	cam.transform.SetPosition(mgl32.Vec3{5, 3, 5})
	cam.transform.LookAt(mgl32.Vec3{0, 0, 0})
	cam.Draw(objShader)

	cube = createCube()

	return true
}

func shutDown() {
	util.LogGLErrors()
	cube.destroy()
	glfw.Terminate()
}


func run() {

	window := glfw.GetCurrentContext()

	angle := 0.0
	previousTime := glfw.GetTime()

	for !window.ShouldClose() {
		// Update
		time := glfw.GetTime()
		elapsed := time - previousTime
		previousTime = time

		angle += elapsed
		cube.transform.objMatrix = mgl32.HomogRotate3D(float32(angle), mgl32.Vec3{0, 1, 0})

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);

		if objShader.Bind() {
			texture.Bind(gl.TEXTURE0)
			sceneLight.Draw(objShader)
			cube.draw(objShader)
		}

		window.SwapBuffers()
		glfw.PollEvents()

		// Check for errors
		util.LogGLErrors()
	}
}
