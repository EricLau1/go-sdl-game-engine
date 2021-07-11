package physics

type Transform struct {
	X, Y float64
}

func NewTransform(x, y float64) *Transform {
	return &Transform{x, y}
}

func (t *Transform) TranslateX(x float64) {
	t.X += x
}

func (t *Transform) TranslateY(y float64) {
	t.Y += y
}

func (t *Transform) Translate(v Vector2D) {
	t.X += v.X
	t.Y += v.Y
}
