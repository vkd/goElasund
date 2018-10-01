package core

import "goElasund/core/wp"

type PlayerColor int

const (
	PlayerColor_Red PlayerColor = iota
	PlayerColor_Green
	PlayerColor_Blue
	PlayerColor_Yellow
)

var (
	// AllPlayers - list of all players
	AllPlayers = []PlayerColor{PlayerColor_Red, PlayerColor_Green, PlayerColor_Blue, PlayerColor_Yellow}
)

// Player of game
type Player struct {
	wp wp.WictoryPoints
}
