package engine

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/core"
	"go-sdl-game-engine/engine/timer"
)

func Run() {
	if core.DefaultEngine.Init() {
		start := sdl.GetTicks()
		var frames int64
		for core.DefaultEngine.IsRunning() {
			frames++
			core.DefaultEngine.Events()
			core.DefaultEngine.Update()
			core.DefaultEngine.Render()
			timer.TickWithFPS(start, frames)
		}
		core.DefaultEngine.Clean()
	}
}
