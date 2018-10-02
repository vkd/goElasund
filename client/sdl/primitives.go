package sdl

// Primitiver - interface of game primitives
type Primitiver interface {
	Top() int32
	Bottom() int32
	Left() int32
	Right() int32
}

// IsIntersect with two primitives
func IsIntersect(p1, p2 Primitiver) bool {
	if p1.Left() > p2.Right() {
		return false
	}
	if p1.Right() < p2.Left() {
		return false
	}
	if p1.Top() > p2.Bottom() {
		return false
	}
	if p1.Bottom() < p2.Top() {
		return false
	}
	return true
}

// Point on game layout
type Point struct {
	X, Y int32
}

var _ Primitiver = (*Point)(nil)
var _ Primitiver = Point{}

// Left border of point
func (p Point) Left() int32 {
	return p.X
}

// Right border of point
func (p Point) Right() int32 {
	return p.X
}

// Top border of point
func (p Point) Top() int32 {
	return p.Y
}

// Bottom border of point
func (p Point) Bottom() int32 {
	return p.Y
}

// Rect - game rectange
type Rect struct {
	X, Y, W, H int32
}

var _ Primitiver = (*Rect)(nil)
var _ Primitiver = Rect{}

// Left border of rectangle
func (r Rect) Left() int32 {
	return r.X
}

// Right border of rectangle
func (r Rect) Right() int32 {
	return r.X + r.W
}

// Top border of rectangle
func (r Rect) Top() int32 {
	return r.Y
}

// Bottom border of rectangle
func (r Rect) Bottom() int32 {
	return r.Y + r.H
}
