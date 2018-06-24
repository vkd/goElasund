package dice

import (
	"testing"
)

func testManyRolling(t *testing.T, d Dicer, count int, expect map[int]int) {
	result := rollingDicer(d, count)
	if len(result) != len(expect) {
		t.Errorf("Wrong len result of stack dices: %d", len(result))
	}
	for k, v := range expect {
		if result[k] != v {
			t.Errorf("Wrong expect value on sum %d, expect: %d, but get: %d", k, v, result[k])
		}
	}
}

func rollingDicer(d Dicer, count int) map[int]int {
	res := make(map[int]int)
	for i := 0; i < count; i++ {
		res[d.Roll()]++
	}
	return res
}
