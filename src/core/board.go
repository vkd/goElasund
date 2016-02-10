package core

type Board struct {
	Cells [][]Tiler

	width, height int
}

func NewBoard(width, height int) (b *Board) {
	b = new(Board)
	b.width = width
	b.height = height
	b.Cells = make([][]Tiler, width)
	for i, _ := range b.Cells {
		b.Cells[i] = make([]Tiler, height)
	}
	return b
}

func (b *Board) PlaceBuilding(building *Building, x, y int) {
	if !b.checkFreeBuilding(building, x, y, color) {
		return
	}

	b.Cells[x][y] = building

	for i := 0; i < building.Width; i++ {
		for j := 0; j < building.Height; j++ {
			if i == j && i == 0 {
				continue
			}
			b.Cells[x+i][y+j] = &Ref{Object: building}
		}
	}
}

func (b *Board) checkFreeBuilding(building *Building, x, y int, color PlayerColor) bool {
	if building.IsBuild {
		return false
	}
	for i := x; i < x+building.Width; i++ {
		if i >= b.width {
			return false
		}
		for j := y; j < y+building.Height; j++ {
			if j >= b.height {
				return false
			}
			if b.Cells[x][y] != nil {
				return false
			}
		}
	}
	return true
}
