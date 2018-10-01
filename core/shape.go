package core

// Shape of building
type Shape [][]bool

// Shapes
var (
	Shape1x1 = Shape{{true}}
	Shape1x2 = Shape{{true, true}}
	Shape2x1 = Shape{{true}, {true}}
	Shape2x2 = Shape{{true, true}, {true, true}}
	Shape3x2 = Shape{{true, true}, {true, true}, {true, true}}
)
