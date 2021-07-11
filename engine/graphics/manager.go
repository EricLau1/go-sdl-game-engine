package graphics

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureManager interface {
	Load(id, filename string) bool
	DrawFrame(id string, x, y, width, height, row, frame int32, flip sdl.RendererFlip)
	DrawTile(id string, tileSize, x, y, row, frame int32, flip sdl.RendererFlip)
	Drop(id string)
	Clean()
}

type textureManager struct {
	textures map[string]*sdl.Texture
	renderer *sdl.Renderer
}

func NewTextureManager(renderer *sdl.Renderer) TextureManager {
	return &textureManager{
		textures: make(map[string]*sdl.Texture),
		renderer: renderer,
	}
}

func (tm *textureManager) Load(id, filename string) bool {
	surface, err := img.Load(filename)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	defer surface.Free()
	texture, err := tm.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	tm.textures[id] = texture
	return true
}

func (tm *textureManager) DrawFrame(id string, x, y, width, height, row , frame int32, flip sdl.RendererFlip) {
	if texture, exists := tm.textures[id]; exists {

		src := sdl.Rect{X: width * frame, Y: height * row, W: width, H: height}
		dst := sdl.Rect{X: x, Y: y, W: width, H: height}

		err := tm.renderer.CopyEx(texture, &src, &dst, 0, nil, flip)
		if err != nil {
			sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		}
	}
}

func (tm *textureManager) DrawTile(id string, tileSize, x, y, row, frame int32, flip sdl.RendererFlip) {
	if texture, exists := tm.textures[id]; exists {

		src := sdl.Rect{X: tileSize * frame, Y: tileSize * row, W: tileSize, H: tileSize}
		dst := sdl.Rect{X: x, Y: y, W: tileSize, H: tileSize}

		err := tm.renderer.CopyEx(texture, &src, &dst, 0, nil, flip)
		if err != nil {
			sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		}
	}
}

func (tm *textureManager) Drop(id string) {
	err := tm.textures[id].Destroy()
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
	} else {
		delete(tm.textures, id)
	}
}

func (tm *textureManager) Clean() {
	for id, texture := range tm.textures {
		err := texture.Destroy()
		if err != nil {
			sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		} else {
			delete(tm.textures, id)
		}
	}
}
