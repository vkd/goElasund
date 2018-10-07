package sdl

// Button on layout
type Button struct {
	Rect
	ButtonState ButtonState
	Text        string

	fn func()
}

var _ Drawer = (*Button)(nil)

// NewButton - create new button
func NewButton(r Rect, text string, fn func()) *Button {
	return &Button{Rect: r, Text: text, fn: fn}
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

// Draw button
func (b *Button) Draw(draw *Draw) {
	switch b.ButtonState {
	case ButtonStateActive:
		draw.RectByFill(b.Rect, 6, green, orange)
	case ButtonStateInactive:
		return
	case ButtonStateSelected:
		draw.RectByFill(b.Rect, 6, green, blue)
	case ButtonStatePressed:
		draw.RectByFill(b.Rect, 6, green, black)
	}
	draw.TextCenter(b.Text, b.Rect.Center(), yellow, 24)
}

// Update button stage
func (b *Button) Update(e Eventer) {
	switch e := e.(type) {
	case *MouseEvent:
		if IsIntersect(b.Rect, e.Point) {
			b.ButtonState = ButtonStateSelected
			switch e.LMouseState {
			case Pressed:
				b.ButtonState = ButtonStatePressed
			case Released:
				b.ButtonState = ButtonStateActive
				b.fn()
			default:
				b.ButtonState = ButtonStateSelected
			}
		} else {
			b.ButtonState = ButtonStateActive
		}
	}
}
