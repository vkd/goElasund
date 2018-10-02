package sdl

type Layout struct {
	items []*Button
}

func (l *Layout) AddItem(i *Button) {
	l.items = append(l.items, i)
}

func (l *Layout) Draw(d *Draw) {

}
