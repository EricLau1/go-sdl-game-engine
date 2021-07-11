package characters

import (
	"github.com/veandco/go-sdl2/sdl"
	"go-sdl-game-engine/engine/animations"
	"go-sdl-game-engine/engine/collisions"
	"go-sdl-game-engine/engine/graphics"
	"go-sdl-game-engine/engine/inputs"
	"go-sdl-game-engine/engine/physics"
)

const (
	JUMP_TIME   = 15.0
	JUMP_FORCE  = 10.0
	RUN_FORCE   = 4.0
	ATTACK_TIME = 20.0

	GROUND = 200
)

var (
	WarriorIdleFrames = animations.Frames{
		TextureID: "player_idle",
		Count:     8,
		Row:       0,
		Speed:     80,
	}

	WarriorRunFrames = animations.Frames{
		TextureID: "player_run",
		Count:     8,
		Row:       0,
		Speed:     80,
	}

	WarriorJumpFrames = animations.Frames{
		TextureID: "player_jump",
		Count:     2,
		Row:       0,
		Speed:     200,
	}

	WarriorFallFrames = animations.Frames{
		TextureID: "player_fall",
		Count:     2,
		Row:       0,
		Speed:     200,
	}

	WarriorAttack1Frames = animations.Frames{
		TextureID: "player_attack1",
		Count:     4,
		Row:       0,
		Speed:     200,
	}

	WarriorAttack2Frames = animations.Frames{
		TextureID: "player_attack2",
		Count:     4,
		Row:       0,
		Speed:     200,
	}

	WarriorAttack3Frames = animations.Frames{
		TextureID: "player_attack3",
		Count:     4,
		Row:       0,
		Speed:     200,
	}
)

var DefaultWarriorProps = Properties{
	TextureID: "warrior",
	Width:     1280 / 8,
	Height:    111,
	Transform: physics.NewTransform(100, 200),
	Flip:      sdl.FLIP_NONE,
}

type Warrior struct {
	*Character
	isRunning         bool
	jumpTime          float64
	jumpForce         float64
	attackTime        float64
	isJumping         bool
	isFalling         bool
	isGrounded        bool
	isAttacking1      bool
	isAttacking2      bool
	isAttacking3      bool
	rigidBody         *physics.RigidBody
	collider          *collisions.Collider
	lastSafePosition  *physics.Vector2D
	textureManager    graphics.TextureManager
	animationManager  animations.AnimationManager
	collisionsHandler *collisions.CollisionHandler
}

func NewWarrior(props *Properties, textureManager graphics.TextureManager, collisionsHandler *collisions.CollisionHandler) *Warrior {
	var warrior Warrior

	warrior.Character = &Character{"warrior", props}
	warrior.textureManager = textureManager
	warrior.lastSafePosition = physics.NewVector2D(props.Transform.X, props.Transform.Y)

	warrior.props.Flip = sdl.FLIP_NONE
	warrior.jumpTime = JUMP_TIME
	warrior.jumpForce = JUMP_FORCE
	warrior.attackTime = ATTACK_TIME

	warrior.collider = &collisions.Collider{}
	warrior.collider.SetBuffer(-65, -53, 0, 0)

	warrior.rigidBody = physics.NewRigidBody()
	warrior.rigidBody.SetGravity(3.0)

	warrior.animationManager = animations.NewAnimationManager(textureManager)
	warrior.animationManager.Set(&WarriorIdleFrames)

	warrior.collisionsHandler = collisionsHandler

	return &warrior
}

func (w *Warrior) Draw() {
	w.animationManager.Draw(int32(w.GetX()), int32(w.GetY()), w.props.Width, w.props.Height, w.props.Flip)
}

func (w *Warrior) Update(dt float64) {
	w.isRunning = false
	w.rigidBody.UnSetForce()

	if inputs.GetAxisDirection(inputs.HORIZONTAL) == physics.FORWARD && !w.IsAttacking() {
		w.rigidBody.ApplyForceX(physics.FORWARD * RUN_FORCE)
		w.props.Flip = sdl.FLIP_NONE
		w.isRunning = true
	}

	if inputs.GetAxisDirection(inputs.HORIZONTAL) == physics.BACKWARD && !w.IsAttacking() {
		w.rigidBody.ApplyForceX(physics.BACKWARD * RUN_FORCE)
		w.props.Flip = sdl.FLIP_HORIZONTAL
		w.isRunning = true
	}

	if inputs.GetMouseButtonDown(inputs.MOUSE_BUTTON_LEFT) {
		w.rigidBody.UnSetForce()
		w.Attack1()
	}

	if inputs.GetMouseButtonDown(inputs.MOUSE_BUTTON_RIGHT) {
		w.rigidBody.UnSetForce()
		w.Attack2()
	}

	if inputs.GetKeyDown(sdl.SCANCODE_X) {
		w.rigidBody.UnSetForce()
		w.Attack3()
	}

	if inputs.GetKeyDown(sdl.SCANCODE_SPACE) && w.isGrounded {
		w.isJumping = true
		w.isGrounded = false
		w.rigidBody.ApplyForceY(physics.UPWARD * w.jumpForce)
	}

	if inputs.GetKeyDown(sdl.SCANCODE_SPACE) && w.isJumping && w.jumpTime > 0 {
		w.jumpTime -= dt
		w.rigidBody.ApplyForceY(physics.UPWARD * w.jumpForce)
	} else {
		w.isJumping = false
		w.jumpTime = JUMP_FORCE
	}

	if w.rigidBody.Velocity().Y > 0 && !w.isGrounded {
		w.isFalling = true
	} else {
		w.isFalling = false
	}

	if w.IsAttacking() && w.attackTime > 0 {
		w.attackTime -= dt
	} else {
		w.StopAttack()
		w.attackTime = ATTACK_TIME
	}

	boxWidth := int32(30)
	boxHeight := int32(51)

	w.rigidBody.Update(dt)

	w.lastSafePosition.X = w.GetX()
	w.props.Transform.X += w.rigidBody.Position().X
	w.collider.Set(int32(w.GetX()), int32(w.GetY()), boxWidth, boxHeight)

	if w.collisionsHandler.Map(w.collider.Get()) {
		sdl.Log("COLLIDE X !!!")
		w.props.Transform.X = w.lastSafePosition.X
	}

	w.rigidBody.Update(dt)
	w.lastSafePosition.Y = w.GetY()
	w.props.Transform.Y += w.rigidBody.Position().Y
	w.collider.Set(int32(w.GetX()), int32(w.GetY()), boxWidth, boxHeight)

	if w.collisionsHandler.Map(w.collider.Get()) {
		sdl.Log("COLLIDE Y !!!")
		w.isGrounded = true
		w.props.Transform.Y = w.lastSafePosition.Y
	} else {
		w.isGrounded = false
	}

	w.Animate()
	w.animationManager.Update()
}

func (w *Warrior) Clean() {
	w.textureManager.Drop(w.props.TextureID)
}

func (w *Warrior) GetX() float64 {
	return w.props.Transform.X
}

func (w *Warrior) GetY() float64 {
	return w.props.Transform.Y
}

func (w *Warrior) Animate() {
	w.animationManager.Set(&WarriorIdleFrames)

	if w.isRunning {
		w.animationManager.Set(&WarriorRunFrames)
	}

	if w.isJumping {
		w.animationManager.Set(&WarriorJumpFrames)
	}

	if w.isFalling {
		w.animationManager.Set(&WarriorFallFrames)
	}

	if w.isAttacking1 {
		w.animationManager.Set(&WarriorAttack1Frames)
	}

	if w.isAttacking2 {
		w.animationManager.Set(&WarriorAttack2Frames)
	}

	if w.isAttacking3 {
		w.animationManager.Set(&WarriorAttack3Frames)
	}
}

func (w *Warrior) IsAttacking() bool {
	return w.isAttacking1 || w.isAttacking2 || w.isAttacking3
}

func (w *Warrior) StopAttack() {
	w.isAttacking1 = false
	w.isAttacking2 = false
	w.isAttacking3 = false
}

func (w *Warrior) Attack1() {
	w.isAttacking1 = true
	w.isAttacking2 = false
	w.isAttacking3 = false
}

func (w *Warrior) Attack2() {
	w.isAttacking1 = false
	w.isAttacking2 = true
	w.isAttacking3 = false
}

func (w *Warrior) Attack3() {
	w.isAttacking1 = false
	w.isAttacking2 = false
	w.isAttacking3 = true
}
