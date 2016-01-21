package core

import (
	"math/rand"
	"time"
)

type Elasund struct {
	random *rand.Rand

	count_players int

	Buildings []*Building

	Board *Board
}

func NewElasund() (e *Elasund) {
	e = new(Elasund)
	e.random = rand.New(rand.NewSource(time.Now().UnixNano()))

	return e
}

func (e *Elasund) Build(building_type BuildingType, x, y int, color PlayerColor) {
	for _, b := range e.Buildings {
		if b.IsBuild {
			continue
		}
		if b.Type != building_type {
			continue
		}
		b.Build(x, y, color)
		return
	}
}

func (e *Elasund) CheckBuild(building_type BuildingType, x, y int, color PlayerColor) bool {
	index := -1
	for i, b := range e.Buildings {
		if b.IsBuild {
			continue
		}
		if b.Type != building_type {
			continue
		}
		index = i
	}
	if index == -1 {
		return false
	}

	width := e.Buildings[index].Width
	height := e.Buildings[index].Height

	width_map := (e.count_players * 2) + 1
	height_map := 10

	if x+(width-1) >= width_map {
		return false
	}
	if y+(height-1) >= height_map {
		return false
	}

	e.Buildings[index].X = x
	e.Buildings[index].Y = y

	for _, b := range e.Buildings {
		if !b.IsBuild {
			continue
		}
		if b.Intersect(e.Buildings[index]) {
			return false
		}
	}
	return true
}

func (e *Elasund) Initialize(count_players int) {
	e.count_players = count_players
	e.Board = NewBoard(1+(2*count_players), 10)
	e.Board.Cells[4][5] = &Building{Type: DrawWell}
	// for i := 1; i <= 9; i++ {
	// 	e.Buildings = append(e.Buildings, &Building{Type: Church, Value: i})
	// }
	for i := 0; i < 4; i++ {
		e.Buildings = append(e.Buildings, &Building{Type: DrawWell})
	}
	for i := 0; i < 4; i++ {
		e.Buildings = append(e.Buildings, &Building{Type: Fair, Width: 2, Height: 1})
	}
	// for i := 1; i < 3; i++ {
	// 	e.Buildings = append(e.Buildings, &Building{Type: Government, Value: i})
	// }
	for i := 0; i < 5; i++ {
		e.Buildings = append(e.Buildings, &Building{Type: Hotel, Width: 2, Height: 2})
	}
	// for i := 0; i < 5; i++ {
	//     e.Buildings = append(e.Buildings, &Building{Type: House})
	// }

	for i := 0; i < 5; i++ {
		e.Buildings = append(e.Buildings, &Building{Type: Shop})
	}
}
