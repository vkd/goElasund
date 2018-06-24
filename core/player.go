package core

import "goElasund/core/wp"

type PlayerColor int

const (
	PlayerColor_Red PlayerColor = iota
	PlayerColor_Blue
	PlayerColor_Green
	PlayerColor_Yellow
)

// Player of game
type Player struct {
	wp wp.WictoryPoints
}
