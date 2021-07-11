package collisions

import "github.com/veandco/go-sdl2/sdl"

type Collider struct {
	box *sdl.Rect
	buf *sdl.Rect
}

func (c *Collider) Set(x, y, w, h int32) {
	c.box = &sdl.Rect{
		X: x - c.buf.X,
		Y: y - c.buf.Y,
		W: w - c.buf.W,
		H: h - c.buf.H,
	}
}

func (c *Collider) Get() sdl.Rect {
	return *c.box
}

func (c *Collider) SetBuffer(x, y, w, h int32) {
	c.buf = &sdl.Rect{X: x, Y: y, W: w, H: h}
}


