package point

import "goElasund/core"

type Point core.Vector

func (p *Point) Add(point Point) *Point {
	p.X += point.X
	p.Y += point.Y
	return p
}

func (p *Point) Move(x int, y int) *Point {
	p.X += x
	p.Y += y
	return p
}
