package wp

const (
	maxWP = 10
)

// WictoryPoints of player
type WictoryPoints struct {
	wp int
}

// Add wp
func (w *WictoryPoints) Add(count int) {
	w.wp += count
	if w.wp < 0 {
		w.wp = 0
	}
}

// IsWinner - check wps is enough to win the game
func (w *WictoryPoints) IsWinner() bool {
	return w.wp >= maxWP
}
