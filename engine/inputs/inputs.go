package inputs

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Axis int

const (
	HORIZONTAL Axis = 0
	VERTICAL   Axis = 1
)

var keyboard = sdl.GetKeyboardState()

func Listen(quit func()) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.GetType() {
		case sdl.QUIT:
			quit()
		case sdl.KEYDOWN:
			KeyDown()
		case sdl.KEYUP:
			KeyUp()
		case sdl.MOUSEBUTTONDOWN:
			OnMouseButtonDown(event.(*sdl.MouseButtonEvent))
		case sdl.MOUSEBUTTONUP:
			OnMouseButtonUp(event.(*sdl.MouseButtonEvent))
		}
	}
}

func GetAxisDirection(axis Axis) int {
	switch axis {
	case HORIZONTAL:
		if GetKeyDown(sdl.SCANCODE_D) || GetKeyDown(sdl.SCANCODE_RIGHT) {
			return 1
		}
		if GetKeyDown(sdl.SCANCODE_A) || GetKeyDown(sdl.SCANCODE_LEFT) {
			return -1
		}
	case VERTICAL:
		if GetKeyDown(sdl.SCANCODE_W) || GetKeyDown(sdl.SCANCODE_UP) {
			return 1
		}
		if GetKeyDown(sdl.SCANCODE_S) || GetKeyDown(sdl.SCANCODE_DOWN) {
			return -1
		}
	}
	return 0
}

func KeyDown() {
	keyboard = sdl.GetKeyboardState()
}

func KeyUp() {
	keyboard = sdl.GetKeyboardState()
}

func GetKeyDown(key sdl.Scancode) bool {
	return keyboard[key] == 1
}
