package camera

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/physics"
)

type Camera struct {
	target       *physics.Point
	position     *physics.Vector2D
	viewBox      sdl.Rect
	screenWidth  int32
	screenHeight int32
}

func NewCamera(screenWidth, screenHeight int32) *Camera {
	return &Camera{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		viewBox: sdl.Rect{
			X: 0,
			Y: 0,
			W: screenWidth,
			H: screenHeight,
		},
	}
}

func (c *Camera) SetTarget(point *physics.Point) {
	c.target = point
}

func (c *Camera) Update(dt float64) {
	if c.target != nil {

		c.viewBox.X = int32(c.target.X) - c.screenWidth/2
		c.viewBox.Y = int32(c.target.Y) - c.screenHeight/2

		if c.viewBox.X < 0 {
			c.viewBox.X = 0
		}

		if c.viewBox.Y < 0 {
			c.viewBox.Y = 0
		}

		width := 2*c.screenWidth - c.viewBox.W
		if c.viewBox.X > width {
			c.viewBox.X = width
		}

		height := c.screenHeight - c.viewBox.H
		if c.viewBox.Y > height {
			c.viewBox.Y = height
		}

		c.position = physics.NewVector2D(float64(c.viewBox.X), float64(c.viewBox.Y))
	}
}

func (c *Camera) ViewBox() sdl.Rect {
	return c.viewBox
}

func (c *Camera) Position() physics.Vector2D {
	return *c.position
}

func (c *Camera) Move(scalar float64) physics.Vector2D {
	return c.position.Mul(scalar)
}
