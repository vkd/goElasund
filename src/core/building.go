package core

type BuildingType int

const (
	BuildingType_Church BuildingType = iota
	BuildingType_DrawWell
	BuildingType_Fair
	BuildingType_Government
	BuildingType_Hotel
	BuildingType_House
	BuildingType_Shop
	BuildingType_SmallTotem
	BuildingType_Totem
	BuildingType_Workshop
)

type Building struct {
	X, Y          int
	Width, Height int

	IsBuild bool

	Type  BuildingType
	Color PlayerColor

	Value int
}

func (b *Building) GetType() TileType {
	return TileType_Building
}

func (b *Building) Build(x, y int, color PlayerColor) {
	b.X = x
	b.Y = y
	b.Color = color
	b.IsBuild = true
}

func (b *Building) Destroy() {
	b.IsBuild = false
}

func (b *Building) Right() int {
	return b.X + (b.Width - 1)
}

func (b *Building) Bottom() int {
	return b.Y + (b.Height - 1)
}

func (b *Building) Intersect(target *Building) bool {
	if b.Right() < target.X {
		return false
	}
	if b.X > target.Right() {
		return false
	}
	if b.Bottom() < target.Y {
		return false
	}
	if b.Y > target.Bottom() {
		return false
	}
	return true
}
