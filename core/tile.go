package core

type Tile struct {
	X, Y int
}

func (t *Tile) Place(x int, y int) {
	t.X = x
	t.Y = y
}
