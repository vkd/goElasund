package wp

import (
	"testing"
)

func TestWictoryPoints_Add(t *testing.T) {
	tests := []struct {
		name  string
		start int
		want  bool
	}{
		{"0", 0, false},
		{"1", 1, false},
		{"2", 2, false},
		{"3", 3, false},
		{"4", 4, false},
		{"5", 5, false},
		{"6", 6, false},
		{"7", 7, false},
		{"8", 8, false},
		{"9", 9, true},
		{"10", 10, true},
		{"11", 11, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &WictoryPoints{
				wp: tt.start,
			}
			w.Add(1)
			if w.IsWinner() != tt.want {
				t.Errorf("Error in checker winner player: %v (want: %v)", w.IsWinner(), tt.want)
			}
		})
	}
}

func TestWictoryPoints_IsWinner(t *testing.T) {
	var wp WictoryPoints
	for i := 0; i < 9; i++ {
		wp.Add(1)
		if wp.IsWinner() {
			t.Fatalf("Wrong in isWinner checker: %d", i)
		}
	}
	for i := 0; i < 3; i++ {
		wp.Add(1)
		if !wp.IsWinner() {
			t.Fatalf("Wrong in positive isWinner: %d", i)
		}
	}
}
