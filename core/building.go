package core

// BuildingType - type of buildings
type BuildingType int

// Buildings
const (
	Church BuildingType = iota
	DrawWell
	Fair
	Government
	Hotel
	House
	Shop
	SmallTotem
	Totem
	Workshop
)

// Building of game
type Building struct {
	Tile

	IsBuilded bool
	Type      BuildingType
	Index     int

	Color PlayerColor
}

var _ tiler = (*Building)(nil)

// func (b *Building) Build(x, y int, color PlayerColor) {
// 	b.X = x
// 	b.Y = y
// 	b.Color = color
// 	b.IsBuilded = true
// }

func (b *Building) Destroy() {
	b.IsBuilded = false
}

// func (b *Building) Right() int {
// 	return b.X + (b.Width - 1)
// }

// func (b *Building) Bottom() int {
// 	return b.Y + (b.Height - 1)
// }

// func (b *Building) Intersect(target *Building) bool {
// 	if b.Right() < target.X {
// 		return false
// 	}
// 	if b.X > target.Right() {
// 		return false
// 	}
// 	if b.Bottom() < target.Y {
// 		return false
// 	}
// 	if b.Y > target.Bottom() {
// 		return false
// 	}
// 	return true
// }
