package physics

type Vector2D struct {
	X float64
	Y float64
}

func NewVector2D(x, y float64) *Vector2D {
	return &Vector2D{x, y}
}

func (v2d *Vector2D) Sum(other Vector2D) Vector2D {
	return Vector2D{
		X: v2d.X + other.X,
		Y: v2d.Y + other.Y,
	}
}

func (v2d *Vector2D) Sub(other Vector2D) Vector2D {
	return Vector2D{
		X: v2d.X - other.X,
		Y: v2d.Y - other.Y,
	}
}

func (v2d *Vector2D) Mul(scalar float64) Vector2D {
	return Vector2D{
		X: v2d.X * scalar,
		Y: v2d.Y * scalar,
	}
}
