package core

type BuildingType int

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

type Building struct {
	X, Y int

	OnMap bool

	Type  BuildingType
	Color PlayerColor

	Value int
}

func (b *Building) Build(x, y int, color PlayerColor) {
	b.X = x
	b.Y = y
	b.Color = color
	b.OnMap = true
}

func (b *Building) Destroy() {
	b.OnMap = false
}
