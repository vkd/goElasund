package core

type Ref struct {
	Object Tiler
}

func (r *Ref) GetType() TileType {
	return TileType_Ref
}
