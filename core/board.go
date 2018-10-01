package core

type Board struct {
	// Cells [][]Tile

	// width, height int

	Buildings []*Building
	Claims    []*Building
}

func NewBoard(width, height int) (b *Board) {
	b = new(Board)
	// b.width = width
	// b.height = height
	// b.Cells = make([][]Tile, width)
	// for i, _ := range b.Cells {
	// 	b.Cells[i] = make([]Tile, height)
	// }
	return b
}

type tiler interface {
	IsIntersect(t Tile) bool
}

func (b *Board) GetIntersections(t Tile) []tiler {
	var out []tiler
	for i, bs := range b.Buildings {
		if bs.IsIntersect(t) {
			out = append(out, b.Buildings[i])
		}
	}
	return out
}

// func (b *Board) PlaceBuilding(building *Building, x, y int) {
// if !b.checkFreeBuilding(building, x, y, color) {
// 	return
// }

// b.Cells[x][y] = building

// for i := 0; i < building.Width; i++ {
// 	for j := 0; j < building.Height; j++ {
// 		if i == j && i == 0 {
// 			continue
// 		}
// 		b.Cells[x+i][y+j] = &Ref{Object: building}
// 	}
// }
// }

// func (b *Board) checkFreeBuilding(building *Building, x, y int, color PlayerColor) bool {
// 	if building.IsBuild {
// 		return false
// 	}
// 	for i := x; i < x+building.Width; i++ {
// 		if i >= b.width {
// 			return false
// 		}
// 		for j := y; j < y+building.Height; j++ {
// 			if j >= b.height {
// 				return false
// 			}
// 			if b.Cells[x][y] != nil {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
