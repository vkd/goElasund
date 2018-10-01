package core

import "testing"

func TestTile_IsIntersect(t *testing.T) {
	tests := []struct {
		name string
		a    Tile
		b    Tile
		want bool
	}{
		// TODO: Add test cases.
		{"base ok", Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, true},
		{"error 1", Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, Tile{Vector: Vector{4, 4}, Shape: Shape1x1}, false},
		{"error 2", Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, Tile{Vector: Vector{5, 4}, Shape: Shape1x1}, false},
		{"error 3", Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, Tile{Vector: Vector{6, 4}, Shape: Shape1x1}, false},
		{"error 4", Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, Tile{Vector: Vector{4, 5}, Shape: Shape1x1}, false},
		{"error 5", Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, Tile{Vector: Vector{6, 5}, Shape: Shape1x1}, false},
		{"error 6", Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, Tile{Vector: Vector{4, 6}, Shape: Shape1x1}, false},
		{"error 7", Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, Tile{Vector: Vector{5, 6}, Shape: Shape1x1}, false},
		{"error 8", Tile{Vector: Vector{5, 5}, Shape: Shape1x1}, Tile{Vector: Vector{6, 6}, Shape: Shape1x1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.IsIntersect(tt.b); got != tt.want {
				t.Errorf("Tile.IsIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
