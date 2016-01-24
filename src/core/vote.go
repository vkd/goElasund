package core

import (
	"math/rand"
)

type ColorVote int

const (
	VoteBlue ColorVote = iota
	VoteGreen
	VoteRed
)

type Vote struct {
	Color ColorVote
}

func NewVote() (v *Vote) {
	v = new(Vote)
	v.Color = ColorVote(rand.Int() % 3)
	return v
}
