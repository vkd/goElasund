package dice

import (
	"testing"
)

func TestStackDices_Roll(t *testing.T) {
	expect := map[int]int{
		2:  1,
		3:  2,
		4:  3,
		5:  4,
		6:  5,
		7:  6,
		8:  5,
		9:  4,
		10: 3,
		11: 2,
		12: 1,
	}
	sd := NewStackDices(NewDice(6), NewDice(6))
	testManyRolling(t, sd, 36, expect)
	testManyRolling(t, sd, 36, expect)
	testManyRolling(t, sd, 36, expect)
}
