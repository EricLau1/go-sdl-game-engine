package characters

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/physics"
)

type Properties struct {
	TextureID     string
	Width, Height int32
	Transform     *physics.Transform
	Flip          sdl.RendererFlip
}

type Character struct {
	name  string
	props *Properties
}
