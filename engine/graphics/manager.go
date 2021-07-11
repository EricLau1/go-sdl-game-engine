package graphics

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/camera"
)

type TextureManager interface {
	Load(id, filename string) bool
	Draw(id string, x, y, width, height int32, flip sdl.RendererFlip)
	DrawFrame(id string, x, y, width, height, row, frame int32, flip sdl.RendererFlip)
	DrawTile(id string, tileSize, x, y, row, frame int32, flip sdl.RendererFlip)
	Drop(id string)
	Clean()
}

type textureManager struct {
	textures map[string]*sdl.Texture
	renderer *sdl.Renderer
	cam      *camera.Camera
}

func NewTextureManager(renderer *sdl.Renderer, cam *camera.Camera) TextureManager {
	return &textureManager{
		textures: make(map[string]*sdl.Texture),
		renderer: renderer,
		cam:      cam,
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

func (tm *textureManager) Draw(id string, x, y, width, height int32, flip sdl.RendererFlip) {
	if texture, exists := tm.textures[id]; exists {

		src := sdl.Rect{X: 0, Y: 0, W: width, H: height}
		cam := tm.cam.Move(0.5)
		dst := sdl.Rect{X: x - int32(cam.X), Y: y - int32(cam.Y), W: width, H: height}

		err := tm.renderer.CopyEx(texture, &src, &dst, 0, nil, flip)
		if err != nil {
			sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		}
	}
}

func (tm *textureManager) DrawFrame(id string, x, y, width, height, row, frame int32, flip sdl.RendererFlip) {
	if texture, exists := tm.textures[id]; exists {

		src := sdl.Rect{X: width * frame, Y: height * row, W: width, H: height}
		cam := tm.cam.Position()
		dst := sdl.Rect{X: x - int32(cam.X), Y: y - int32(cam.Y), W: width, H: height}

		err := tm.renderer.CopyEx(texture, &src, &dst, 0, nil, flip)
		if err != nil {
			sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		}
	}
}

func (tm *textureManager) DrawTile(id string, tileSize, x, y, row, frame int32, flip sdl.RendererFlip) {
	if texture, exists := tm.textures[id]; exists {

		src := sdl.Rect{X: tileSize * frame, Y: tileSize * row, W: tileSize, H: tileSize}
		cam := tm.cam.Position()
		dst := sdl.Rect{X: x - int32(cam.X), Y: y - int32(cam.Y), W: tileSize, H: tileSize}

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
