package dice

import (
	"math/rand"
	"time"
)

var _ Dicer = (*Dice)(nil)

// Dice - one real dice with dimensions
type Dice struct {
	rnd   *rand.Rand
	sides int
}

// NewDice - create one dice
func NewDice(sides int) Dice {
	if sides <= 0 {
		sides = 6
	}
	return Dice{
		rnd:   rand.New(rand.NewSource(time.Now().UnixNano())),
		sides: sides,
	}
}

// Roll - throw dice
func (d *Dice) Roll() int {
	return d.rnd.Intn(d.sides) + 1
}
