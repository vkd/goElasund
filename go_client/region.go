package go_client

type Region struct{}

func NewRegion(x, y int) Regioner {
	return &Region{}
}

type Regioner interface{}

func (r *Region) Draw(renderer interface{}) {
	panic("not implemented")
	renderer.Draw()
}

func (r *Region) Add(child Regioner) {
	panic("not implemented")
}
