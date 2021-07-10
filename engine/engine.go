package engine

import (
	"go-sdl-game-engine/engine/core"
	"go-sdl-game-engine/engine/timer"
)

func Run() {
	if core.DefaultEngine.Init() {
		for core.DefaultEngine.IsRunning() {
			core.DefaultEngine.Events()
			core.DefaultEngine.Update()
			core.DefaultEngine.Render()
			timer.Tick()
		}
		core.DefaultEngine.Clean()
	}
}
