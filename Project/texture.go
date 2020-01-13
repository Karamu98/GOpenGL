package main

import (
	"image"
	"os"

	log "./Utilities/Logger"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// Texture ... An OpenGL texture
type Texture struct {
	textureID         uint32
	filePath          string
	textureDimentions mgl32.Vec2
	isValid           bool
}

// Bind ... Binds a texture to the given slot. Returns success
func (texture *Texture) Bind(textureSlot uint32) bool {
	if texture.isValid == true {
		gl.ActiveTexture(textureSlot)
		gl.BindTexture(gl.TEXTURE_2D, texture.textureID)
		return true
	}

	return false
}

// Destroy ... Deletes the texture from OpenGL memory
func (texture *Texture) Destroy() {
	gl.DeleteTextures(1, &texture.textureID)
}

// CreateTexture ... Loads a texture into OpenGL
func CreateTexture(filePath string) *Texture {
	newTex := &Texture{}
	newTex.isValid = false

	// Load image
	imgFile, err := os.Open(filePath)
	if err != nil {
		log.Errorf("Couldn't load texture at %v.", filePath)
		return newTex
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Errorf("Couldn't decode texture at %v. %v", filePath, err)
		return newTex
	}

	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		log.Errorf("Unsupported stride for %v", filePath)
		return newTex
	}

	gl.GenTextures(1, &newTex.textureID)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, newTex.textureID)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA,
		int32(rgba.Rect.Size().X),
		int32(rgba.Rect.Size().Y),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(rgba.Pix))

	newTex.isValid = true
	newTex.filePath = filePath
	newTex.textureDimentions = mgl32.Vec2{float32(rgba.Rect.Size().X), float32(rgba.Rect.Size().Y)}

	return newTex
}
