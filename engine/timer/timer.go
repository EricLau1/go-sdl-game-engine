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

	fps float64
)

// Tick calc deltatime
func Tick() {
	tick := float64(sdl.GetTicks())
	deltaTime = (tick - lastTime) * (TARGET_FPS / 1000.0)
	if deltaTime > TARGET_DELTATIME {
		deltaTime = TARGET_DELTATIME
	}
	lastTime = float64(sdl.GetTicks())
}

// TickWithFPS calc deltatime and fps
func TickWithFPS(startLoop uint32, frames int64) {
	Tick()
	elapsedMS := sdl.GetTicks() - startLoop
	if elapsedMS != 0 {
		elapsedSeconds := elapsedMS / 1000
		fps = float64(frames) / float64(elapsedSeconds)
	}
}

func GetFPS() float64 {
	return fps
}

func DeltaTime() float64 {
	return deltaTime
}
