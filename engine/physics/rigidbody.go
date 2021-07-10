package physics

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	UNI_MASS = 1.0
	GRAVITY  = sdl.STANDARD_GRAVITY
	FORWARD  = 1
	BACKWARD = -1
	UPWARD   = -1
	DOWNWARD = 1
)

type RigidBody struct {
	mass         float64
	gravity      float64
	force        Vector2D
	friction     Vector2D
	position     Vector2D
	velocity     Vector2D
	acceleration Vector2D
}

func NewRigidBody() *RigidBody {
	return &RigidBody{
		mass:    UNI_MASS,
		gravity: GRAVITY,
	}
}

func (rb *RigidBody) SetMass(mass float64) {
	rb.mass = mass
}

func (rb *RigidBody) SetGravity(gravity float64) {
	rb.gravity = gravity
}

func (rb *RigidBody) ApplyForce(force Vector2D) {
	rb.force = force
}

func (rb *RigidBody) ApplyForceX(forceX float64) {
	rb.force.X = forceX
}

func (rb *RigidBody) ApplyForceY(forceY float64) {
	rb.force.Y = forceY
}

func (rb *RigidBody) UnSetForce() {
	rb.force.X, rb.force.Y = 0, 0
}

func (rb *RigidBody) ApplyFriction(friction Vector2D) {
	rb.friction = friction
}

func (rb *RigidBody) UnSetFriction() {
	rb.friction.X = 0
	rb.friction.Y = 0
}

func (rb *RigidBody) Mass() float64 {
	return rb.mass
}

func (rb *RigidBody) Position() Vector2D {
	return rb.position
}

func (rb *RigidBody) Velocity() Vector2D {
	return rb.velocity
}

func (rb *RigidBody) Acceleration() Vector2D {
	return rb.acceleration
}

func (rb *RigidBody) Update(dt float64) {
	rb.acceleration.X = (rb.force.X + rb.friction.X) / rb.mass
	rb.acceleration.Y = rb.gravity + rb.force.Y/rb.mass
	rb.velocity = rb.acceleration.Mul(dt)
	rb.position = rb.velocity.Mul(dt)
}
