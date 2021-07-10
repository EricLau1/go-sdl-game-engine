package animations

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/graphics"
)

type AnimationManager interface {
	Set(frames *Frames)
	Draw(x, y, spriteWidth, spriteHeight int32, flip sdl.RendererFlip)
	Update()
}

type animationManager struct {
	speed          int32
	spriteRow      int32
	frameCount     int32
	currentFrame   int32
	textureID      string
	textureManager graphics.TextureManager
}

func NewAnimationManager(textureManager graphics.TextureManager) AnimationManager {
	return &animationManager{
		textureManager: textureManager,
	}
}

func (am *animationManager) Draw(x, y, spriteWidth, spriteHeight int32, flip sdl.RendererFlip) {
	am.textureManager.DrawFrame(am.textureID, x, y, spriteWidth, spriteHeight, am.spriteRow, am.currentFrame, flip)
}

func (am *animationManager) Update() {
	am.currentFrame = (int32(sdl.GetTicks()) / am.speed) % am.frameCount
}

func (am *animationManager) Set(frames *Frames) {
	am.textureID = frames.TextureID
	am.spriteRow = frames.Row
	am.frameCount = frames.Count
	am.speed = frames.Speed
}
