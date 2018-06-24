package dice

import (
	"math/rand"
	"time"
)

// StackDices - stack of sum dices with uniform distributions
type StackDices struct {
	rnd *rand.Rand

	countAvailable int
	values         []diceValue
}

var _ Dicer = (*StackDices)(nil)

// NewStackDices - create StackDices from slice of Dices
func NewStackDices(dices ...Dice) *StackDices {
	sd := &StackDices{
		rnd:    rand.New(rand.NewSource(time.Now().UnixNano())),
		values: makeStack(dices...),
	}
	sd.countAvailable = len(sd.values)
	return sd
}

// Roll - throw StackDices
func (s *StackDices) Roll() int {
	if s.countAvailable <= 0 {
		s.shuffle()
	}
	index := s.rnd.Intn(s.countAvailable)
	for i, v := range s.values {
		if v.isRolled {
			continue
		}
		if index <= 0 {
			s.values[i].isRolled = true
			s.countAvailable--
			return v.value
		}
		index--
	}
	panic("roll override")
}

func (s *StackDices) shuffle() {
	for i := range s.values {
		s.values[i].isRolled = false
	}
	s.countAvailable = len(s.values)
}

type diceValue struct {
	isRolled bool
	value    int
}

func makeStack(ds ...Dice) []diceValue {
	if len(ds) != 2 {
		panic("Not implemented - too many dices")
	}
	var res []diceValue
	for i := 1; i <= ds[0].sides; i++ {
		for j := 1; j <= ds[1].sides; j++ {
			res = append(res, diceValue{value: i + j})
		}
	}
	return res
}
