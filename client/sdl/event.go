package sdl

// Eventer interface of all game events
type Eventer interface {
	Type() Event
}

// Event - empty base game event
type Event uint8

// Event types
const (
	None Event = iota
	Mouse
)

// MouseEvent of move mouse over game interface
type MouseEvent struct {
	Point
	LMouseState MouseState
}

// Type of event
func (m *MouseEvent) Type() Event { return Mouse }

// MouseState - None, Pressed, Released
type MouseState uint8

// Mouse states
const (
	NoneMouseState MouseState = iota
	Pressed
	Released
)
