package utilities

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	 log "./Logger"
)

// LogGLErrors ... Consumes and logs all OpenGL errors
func LogGLErrors() {

	for {
		err := gl.GetError()
		if err == gl.NO_ERROR {
			break
		}

		log.Errorf("GL Error: %v\n", getGLErrorStr(err))

	}
}

func getGLErrorStr(errCode uint32) string {
	switch errCode {
	case gl.NO_ERROR:
		return "No error"
	case gl.INVALID_ENUM:
		return "Invalid enum"
	case gl.INVALID_VALUE:
		return "Invalid value"
	case gl.INVALID_OPERATION:
		return "Invalid operation"
	case gl.STACK_OVERFLOW:
		return "Stack overflow"
	case gl.STACK_UNDERFLOW:
		return "Stack underflow"
	case gl.OUT_OF_MEMORY:
		return "Out of memory"
	default:
		return "Unknown error"
	}
}
