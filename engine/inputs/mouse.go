package inputs

import (
	"github.com/veandco/go-sdl2/sdl"
)

type MOUSE_BUTTON uint8

const (
	MOUSE_BUTTON_LEFT  MOUSE_BUTTON = sdl.BUTTON_LEFT
	MOUSE_BUTTON_RIGHT MOUSE_BUTTON = sdl.BUTTON_RIGHT
)

var mouseDown *sdl.MouseButtonEvent
var mouseUp *sdl.MouseButtonEvent

func OnMouseButtonDown(event *sdl.MouseButtonEvent) {
	if event != nil {
		mouseDown = event
		mouseUp = nil
	}
}

func OnMouseButtonUp(event *sdl.MouseButtonEvent) {
	if event != nil {
		mouseUp = event
		mouseDown = nil
	}
}

func GetMouseButtonDown(button MOUSE_BUTTON) bool {
	if mouseDown == nil {
		return false
	}
	return mouseDown.Button == uint8(button)
}

func GetMouseButtonUp(button MOUSE_BUTTON) bool {
	if mouseUp == nil {
		return false
	}
	return mouseDown.Button == uint8(button)
}
