package sdl

// Button on layout
type Button struct {
	Rect
	ButtonState ButtonState
	fn          func()
}

// NewButton - create new button
func NewButton(r Rect, fn func()) *Button {
	return &Button{Rect: r, fn: fn}
}

// ButtonState - state of button
type ButtonState int

// Button states
const (
	ButtonStateActive ButtonState = iota
	ButtonStateInactive
	ButtonStateSelected
	ButtonStatePressed
)
