package physics

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{X: x, Y: y}
}

func (p *Point) Sum(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p *Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

func (p *Point) Mul(scalar float64) Point {
	return Point{X: p.X * scalar, Y: p.Y * scalar}
}
