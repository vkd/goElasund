package core

type Board struct {
	Cells [][]Tiler
}

func NewBoard(width, height int) (b *Board) {
	b = new(Board)
	b.Cells = make([][]Tiler, width)
	for i, _ := range b.Cells {
		b.Cells[i] = make([]Tiler, height)
	}
	return b
}
