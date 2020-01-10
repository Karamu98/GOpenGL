package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

var shaderCache map[string]*Shader

// Shader ... A shader program
type Shader struct {
	shaderProgram uint32
	isValid       bool
	shaderPath    string
	uniformCache  map[string]int32
}

func createShader(shaderPath string) *Shader {

	// If the shader is already created, return it
	if shaderCache == nil {
		shaderCache = make(map[string]*Shader)
	} else if shaderCache[shaderPath] != nil {
		return shaderCache[shaderPath]
	}

	sources := preProcess(shaderPath)

	if sources == nil {
		log.Fatal("Could not create shader.")
		return nil
	}

	newShader := &Shader{}

	if !compileSources(newShader, sources) {
		log.Fatal("Could not compile sources")
		return nil
	}

	shaderCache[shaderPath] = newShader

	return newShader
}

const typeToken = "#type"

func shaderTypeFromString(typeString string) (uint32, bool) {

	if typeString == "vertex" {
		return gl.VERTEX_SHADER, false
	} else if typeString == "fragment" {
		return gl.FRAGMENT_SHADER, false
	}

	log.Fatalln("Invalid shader type: " + typeString)

	return 0, true
}

func stringFromShaderType(typeInt uint32) (string, bool) {
	if typeInt == gl.VERTEX_SHADER {
		return "vertex", false
	} else if typeInt == gl.FRAGMENT_SHADER {
		return "fragment", false
	}

	log.Fatalln("Invalid shader type: %i", typeInt)

	return "", true
}

func preProcess(shaderPath string) map[uint32]string {

	file, err := os.Open(shaderPath)
	if err != nil {
		log.Println("File: {0} not found.", shaderPath)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	programs := map[uint32]string{}
	var curProgram uint32
	var isErr bool

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) >= len(typeToken) {

			if line[:len(typeToken)] == typeToken {

				tokenDetail := line[len(typeToken)+1:]
				fmt.Println("Found token: " + tokenDetail)

				curProgram, isErr = shaderTypeFromString(tokenDetail)
				if isErr {
					log.Fatalln(curProgram)
					return nil
				}

				continue
			}
		}

		programs[curProgram] += line + "\n"
	}

	return programs
}

func compileSources(shader *Shader, sources map[uint32]string) bool {

	newShaderProgram := gl.CreateProgram()

	var shaderIDs [5]uint32
	var curShaderElement uint8 = 0

	for key, source := range sources {

		sType, isError := stringFromShaderType(key)
		if isError {
			return false
		}

		fmt.Println("Compiling " + sType + " shader.")

		// Create and compile the shader
		shader := gl.CreateShader(key)
		sourceLen := int32(len(source))
		csources, free := gl.Strs(source)
		gl.ShaderSource(shader, 1, csources, &sourceLen)
		free()
		gl.CompileShader(shader)

		// Check for failure
		var status int32
		gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
		if status == gl.FALSE {
			var logLength int32
			gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

			logMessage := strings.Repeat("\x00", int(logLength+1))
			gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logMessage))

			shaderName, validKey := stringFromShaderType(key)
			if !validKey {

				log.Fatalln(shaderName + " failed to compile:\n" + logMessage)
			}

			return false
		}

		// Attach shader to the program
		gl.AttachShader(newShaderProgram, shader)
		shaderIDs[curShaderElement] = shader

		curShaderElement++
	}

	gl.LinkProgram(newShaderProgram)

	var status int32
	gl.GetProgramiv(newShaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(newShaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		logMessage := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(newShaderProgram, logLength, nil, gl.Str(logMessage))

		gl.DeleteProgram(newShaderProgram)

		for i := curShaderElement - 1; i > 0; i-- {
			gl.DeleteShader(shaderIDs[i])
		}

		log.Fatalln("Program failed to link:\n %s", logMessage)

		return false
	}

	for i := curShaderElement - 1; i > 0; i-- {
		gl.DetachShader(newShaderProgram, shaderIDs[i])
		gl.DeleteShader(shaderIDs[i])
	}

	shader.shaderProgram = newShaderProgram
	shader.isValid = true
	shader.uniformCache = make(map[string]int32)
	return true
}

func getUniformLocation(shader *Shader, uniformName string) int32 {

	loc, exists := shader.uniformCache[uniformName]

	if exists {
		return loc
	}

	loc = gl.GetUniformLocation(shader.shaderProgram, gl.Str(uniformName+"\x00"))

	shader.uniformCache[uniformName] = loc
	return loc
}

func setBool(shader *Shader, uniformName string, value bool) {
	var loc int32 = getUniformLocation(shader, uniformName)

	if loc == -1 {
		fmt.Println("Uniform %s not found in shader.", uniformName)
	} else {
		if value {
			gl.Uniform1i(loc, 1)
		} else {
			gl.Uniform1i(loc, 0)
		}

	}
}

func setInt(shader *Shader, uniformName string, value int32) {
	var loc int32 = getUniformLocation(shader, uniformName)

	if loc == -1 {
		fmt.Println("Uniform %s not found in shader.", uniformName)
	} else {
		gl.Uniform1i(loc, value)
	}
}

func setFloat(shader *Shader, uniformName string, value float32) {
	var loc int32 = getUniformLocation(shader, uniformName)

	if loc == -1 {
		fmt.Println("Uniform %s not found in shader.", uniformName)
	} else {
		gl.Uniform1f(loc, value)
	}
}

func setVec2(shader *Shader, uniformName string, value mgl32.Vec2) {
	var loc int32 = getUniformLocation(shader, uniformName)

	if loc == -1 {
		fmt.Println("Uniform %s not found in shader.", uniformName)
	} else {
		gl.Uniform2f(loc, value.X(), value.Y())
	}
}

func setVec3(shader *Shader, uniformName string, value mgl32.Vec3) {
	var loc int32 = getUniformLocation(shader, uniformName)

	if loc == -1 {
		fmt.Println("Uniform %s not found in shader.", uniformName)
	} else {
		gl.Uniform3f(loc, value.X(), value.Y(), value.Z())
	}
}

func setVec4(shader *Shader, uniformName string, value mgl32.Vec4) {
	var loc int32 = getUniformLocation(shader, uniformName)

	if loc == -1 {
		fmt.Println("Uniform %s not found in shader.", uniformName)
	} else {
		gl.Uniform4f(loc, value.X(), value.Y(), value.Z(), value.W())
	}
}

func setMat3(shader *Shader, uniformName string, value mgl32.Mat3) {
	var loc int32 = getUniformLocation(shader, uniformName)

	if loc == -1 {
		fmt.Println("Uniform %s not found in shader.", uniformName)
	} else {
		gl.UniformMatrix3fv(loc, 1, false, &value[0])
	}
}

func setMat4(shader *Shader, uniformName string, value mgl32.Mat4) {
	var loc int32 = getUniformLocation(shader, uniformName)

	if loc == -1 {
		fmt.Println("Uniform %s not found in shader.", uniformName)
	} else {
		gl.UniformMatrix4fv(loc, 1, false, &value[0])
	}
}

func bindShader(shader *Shader) bool {
	if shader.isValid {
		gl.UseProgram(shader.shaderProgram)
		return true
	}

	return false
}

func unbindShader() {
	gl.UseProgram(0)
}