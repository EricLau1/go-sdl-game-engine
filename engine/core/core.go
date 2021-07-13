package core

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/camera"
	"go-sdl-game-engine/engine/characters"
	"go-sdl-game-engine/engine/collisions"
	"go-sdl-game-engine/engine/graphics"
	"go-sdl-game-engine/engine/inputs"
	"go-sdl-game-engine/engine/maps"
	"go-sdl-game-engine/engine/sounds"
	"go-sdl-game-engine/engine/timer"
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
	isRunning        bool
	window           *sdl.Window
	renderer         *sdl.Renderer
	textureManager   graphics.TextureManager
	player           *characters.Warrior
	mapParser        *maps.MapParser
	collisionHandler *collisions.CollisionHandler
	cam              *camera.Camera
	soundsManager    *sounds.Manager
}

func (e *engine) Init() bool {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	sdl.Log("sdl initialized")
	err = img.Init(img.INIT_JPG | img.INIT_PNG)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	sdl.Log("sdl img initialized")
	err = mix.Init(mix.INIT_MP3)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	sdl.Log("sdl mixer initialized")
	err = mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	e.window, err = sdl.CreateWindow(
		"Golang SDL Game Engine v0.0.1",
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		SCREEN_WIDTH, SCREEN_HEIGHT,
		sdl.WINDOW_RESIZABLE|sdl.WINDOW_ALLOW_HIGHDPI,
	)
	sdl.Log("window created")
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	e.renderer, err = sdl.CreateRenderer(e.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	sdl.Log("renderer created")
	e.isRunning = true
	sdl.Log("engined initialized")

	e.cam = camera.NewCamera(SCREEN_WIDTH, SCREEN_HEIGHT)

	textureManager := graphics.NewTextureManager(e.renderer, e.cam)
	e.textureManager = textureManager

	textureManager.Load("player_idle", "assets/player/Idle.png")
	textureManager.Load("player_run", "assets/player/Run.png")
	textureManager.Load("player_jump", "assets/player/Jump.png")
	textureManager.Load("player_fall", "assets/player/Fall.png")
	textureManager.Load("player_attack1", "assets/player/Attack1.png")
	textureManager.Load("player_attack2", "assets/player/Attack2.png")
	textureManager.Load("player_attack3", "assets/player/Attack3.png")

	textureManager.Load("bg", "assets/bg/1.jpg")

	mapParser := maps.NewMapParser("MAP", "assets/maps/map2.tmx", textureManager)
	e.mapParser = mapParser

	e.collisionHandler = collisions.NewCollisionHandler(e.mapParser.GetMap("MAP"))

	e.player = characters.NewWarrior(&characters.DefaultWarriorProps, textureManager, e.collisionHandler)
	e.cam.SetTarget(e.player.GetOrigin())

	e.soundsManager = sounds.NewSoundsManager()
	e.soundsManager.Load("bg", "assets/audios/bg/bg.mp3")
	e.soundsManager.Play("bg", sounds.REPEAT)

	return e.isRunning
}

func (e *engine) Update() {
	dt := timer.DeltaTime()
	e.player.Update(dt)
	e.cam.Update(dt)
	e.window.SetTitle(fmt.Sprintf("FPS: %.1f", timer.GetFPS()))
}

func (e *engine) Render() {
	err := e.renderer.SetDrawColor(124, 218, 254, 255)
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return
	}
	err = e.renderer.Clear()
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return
	}
	e.textureManager.Draw("bg", 0, 0, 2100, SCREEN_HEIGHT+10, sdl.FLIP_NONE)
	e.mapParser.GetMap("MAP").Render()
	e.player.Draw()
	e.renderer.Present()
}

func (e *engine) Events() {
	inputs.Listen(e.Quit)
}

func (e *engine) Clean() bool {
	e.soundsManager.Clean()
	e.mapParser.Clean()
	e.textureManager.Clean()

	err := e.renderer.Destroy()
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	err = e.window.Destroy()
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, err.Error())
		return false
	}
	mix.CloseAudio()
	mix.Quit()
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
