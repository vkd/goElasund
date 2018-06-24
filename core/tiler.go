package core

type Tiler interface {
	GetType() TileType
}

type TileType int

const (
	TileType_Building TileType = iota
	TileType_Claim
	TileType_Ref
)
