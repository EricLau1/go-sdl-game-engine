package core

import (
	"go-sdl-game-engine/engine/characters"
	"go-sdl-game-engine/engine/graphics"
	"go-sdl-game-engine/engine/inputs"
	"go-sdl-game-engine/engine/timer"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_WIDTH  = 960
	SCREEN_HEIGHT = 640
)

var DefaultEngine = &engine{}

type Engine interface {
	Init() bool
	Update()
	Render()
	Events()
	Clean() bool
	IsRunning() bool
}

type engine struct {
	isRunning      bool
	window         *sdl.Window
	renderer       *sdl.Renderer
	textureManager graphics.TextureManager
	player         *characters.Warrior
}

func (e *engine) Init() bool {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_ERROR, err.Error())
		return false
	}
	sdl.Log("sdl initialized")
	err = img.Init(img.INIT_JPG | img.INIT_PNG)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_ERROR, err.Error())
		return false
	}
	sdl.Log("sdl img initialized")
	e.window, err = sdl.CreateWindow(
		"Golang SDL Game Engine v0.0.1",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		SCREEN_WIDTH, SCREEN_HEIGHT,
		sdl.WINDOW_RESIZABLE|sdl.WINDOW_ALLOW_HIGHDPI,
	)
	sdl.Log("window created")
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_ERROR, err.Error())
		return false
	}
	e.renderer, err = sdl.CreateRenderer(e.window, -1, sdl.RENDERER_ACCELERATED | sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_ERROR, err.Error())
		return false
	}
	sdl.Log("renderer created")
	e.isRunning = true
	sdl.Log("engined initialized")

	textureManager := graphics.NewTextureManager(e.renderer)
	e.textureManager = textureManager

	textureManager.Load("player_idle", "assets/player/Idle.png")
	textureManager.Load("player_run", "assets/player/Run.png")
	textureManager.Load("player_jump", "assets/player/Jump.png")
	textureManager.Load("player_fall", "assets/player/Fall.png")
	textureManager.Load("player_attack1", "assets/player/Attack1.png")
	textureManager.Load("player_attack2", "assets/player/Attack2.png")
	textureManager.Load("player_attack3", "assets/player/Attack3.png")

	e.player = characters.NewWarrior(&characters.DefaultWarriorProps, textureManager)

	return e.isRunning
}

func (e *engine) Update() {
	dt := timer.DeltaTime()
	e.player.Update(dt)
}

func (e *engine) Render() {
	err := e.renderer.SetDrawColor(124, 218, 254, 255)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_ERROR, err.Error())
		return
	}
	err = e.renderer.Clear()
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_ERROR, err.Error())
		return
	}
	e.player.Draw()
	e.renderer.Present()
}

func (e *engine) Events() {
	inputs.Listen(e.Quit)
}

func (e *engine) Clean() bool {
	e.textureManager.Clean()

	err := e.renderer.Destroy()
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_ERROR, err.Error())
		return false
	}
	err = e.window.Destroy()
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_ERROR, err.Error())
		return false
	}
	img.Quit()
	sdl.Quit()
	sdl.Log("engine finished")
	return true
}

func (e *engine) IsRunning() bool {
	return e.isRunning
}

func (e *engine) Quit() {
	e.isRunning = false
}
