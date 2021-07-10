package timer

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	TARGET_FPS       = 60
	TARGET_DELTATIME = 1.5
)

var (
	deltaTime float64
	lastTime  float64
)

func Tick() {
	tick := float64(sdl.GetTicks())
	deltaTime = (tick - lastTime) * (TARGET_FPS / 1000.0)
	if deltaTime > 0 {
		//fmt.Printf("CurrentTime=%f, LastTime=%f, DeltaTime=%f\n", tick, lastTime, deltaTime)
	}
	if deltaTime > TARGET_DELTATIME {
		deltaTime = TARGET_DELTATIME
	}
	lastTime = float64(sdl.GetTicks())
}

func DeltaTime() float64 {
	return deltaTime
}
