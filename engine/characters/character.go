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
	name   string
	props  *Properties
	origin *physics.Point
}

func NewCharacter(name string, props *Properties) *Character {
	var c Character
	c.name = name
	c.props = props
	px := props.Transform.X + float64(props.Width)/2.0
	py := props.Transform.Y + float64(props.Height)/2.0
	c.origin = physics.NewPoint(px, py)
	return &c
}

func (c *Character) GetX() float64 {
	return c.props.Transform.X
}

func (c *Character) GetY() float64 {
	return c.props.Transform.Y
}

func (c *Character) GetOrigin() *physics.Point {
	return c.origin
}
