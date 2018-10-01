package core

// Tile is a map placed object
type Tile struct {
	Vector
	Shape Shape
}

var _ tiler = Tile{}
var _ tiler = (*Tile)(nil)

// Top - top border of tile
func (t Tile) Top() int {
	return t.X
}

// Bottom - bottom line of tile
func (t Tile) Bottom() int {
	return t.X + len(t.Shape[0]) - 1
}

// Left - left line of tile
func (t Tile) Left() int {
	return t.Y
}

// Right - right line of tile
func (t Tile) Right() int {
	return t.Y + len(t.Shape) - 1
}

// IsIntersect - check intersection with tile
func (t Tile) IsIntersect(a Tile) bool {
	if a.Top() > t.Bottom() || a.Bottom() < t.Top() {
		return false
	}
	if a.Left() > t.Right() || a.Right() < t.Left() {
		return false
	}

	// TODO: make intersect not only with rectangles
	return true
}
