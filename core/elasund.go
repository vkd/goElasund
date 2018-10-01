package core

import (
	"math/rand"
	"time"
)

// Elasund - a board game
type Elasund struct {
	random *rand.Rand

	countPlayers int

	Buildings []*Building

	Board *Board
}

func NewElasund() (e *Elasund) {
	e = new(Elasund)
	e.random = rand.New(rand.NewSource(time.Now().UnixNano()))

	return e
}

func (e *Elasund) GetBuilding(bt BuildingType, index int) *Building {
	for i, b := range e.Buildings {
		if b.Type == bt && b.Index == index {
			return e.Buildings[i]
		}
	}
	return nil
}

func (e *Elasund) GetBuildingColor(bt BuildingType, color PlayerColor) *Building {
	for i, b := range e.Buildings {
		if b.Type == bt && b.Color == color {
			return e.Buildings[i]
		}
	}
	return nil
}

func (e *Elasund) AddBuilding(b *Building) {
	e.Buildings = append(e.Buildings, b)
}

func (e *Elasund) Build(building_type BuildingType, x, y int, color PlayerColor) {
	// var building *Building
	// for _, b := range e.Buildings {
	// 	if b.IsBuild {
	// 		continue
	// 	}
	// 	if b.Type != building_type {
	// 		continue
	// 	}
	// 	building = b
	// 	break
	// }
	// if b == nil {
	// 	return
	// }
	// b.Build(x, y, color)
	// e.Board.PlaceBuilding(b, x, y)
}

// func (e *Elasund) CheckBuild(building_type BuildingType, x, y int, color PlayerColor) bool {
// 	index := -1
// 	for i, b := range e.Buildings {
// 		if b.IsBuild {
// 			continue
// 		}
// 		if b.Type != building_type {
// 			continue
// 		}
// 		index = i
// 	}
// 	if index == -1 {
// 		return false
// 	}

// 	width := e.Buildings[index].Width
// 	height := e.Buildings[index].Height

// 	width_map := (e.count_players * 2) + 1
// 	height_map := 10

// 	if x+(width-1) >= width_map {
// 		return false
// 	}
// 	if y+(height-1) >= height_map {
// 		return false
// 	}

// 	e.Buildings[index].X = x
// 	e.Buildings[index].Y = y

// 	for _, b := range e.Buildings {
// 		if !b.IsBuild {
// 			continue
// 		}
// 		if b.Intersect(e.Buildings[index]) {
// 			return false
// 		}
// 	}
// 	return true
// }

// Initialize elasund game
func (e *Elasund) Initialize(players []PlayerColor) {
	e.countPlayers = len(players)

	e.Board = NewBoard(1+(2*e.countPlayers), 10)
	// e.Board.Cells[4][5] = newDrawWell()

	// for i := 1; i <= 9; i++ {
	// 	e.Buildings = append(e.Buildings, &Building{Type: Church, Value: i})
	// }

	// Buildings
	// ----------------------------------
	for _, p := range players {
		for bt, v := range totemPlaces[p] {
			e.buildi
			e.AddBuilding(&Building{
				Tile:      Tile{Vector: v},
				IsBuilded: true,
				Type:      bt,
				Color:     p,
				Index:     int(p),
			})
		}
	}
	// e.AddBuilding(&Building{
	// 	Tile: Tile{
	// 		Vector: Vector{3, 2},
	// 	},
	// 	IsBuilded: true,
	// 	Type:      SmallTotem,
	// 	Color:     PlayerColor_Red,
	// })
	// e.AddBuilding(&Building{
	// 	Tile: Tile{
	// 		Vector: Vector{4, 5},
	// 	},
	// 	IsBuilded: true,
	// 	Type:      Totem,
	// 	Color:     PlayerColor_Red,
	// })
	e.Buildings = append(e.Buildings, &Building{
		Tile: Tile{
			Vector: Vector{2, 4},
		},
		IsBuilded: true,
		Type:      Hotel,
	})
	// for i := 0; i < 4; i++ {
	// 	e.Buildings = append(e.Buildings, newDrawWell())
	// }
	// for i := 0; i < 4; i++ {
	// 	e.Buildings = append(e.Buildings, newFair())
	// }
	// for i := 0; i < 3; i++ {
	// 	e.Buildings = append(e.Buildings, newGovernment(i+1))
	// }
	// ----------------------------------

	// b := e.Buildings[5]
	// b.Build(4, 5, PlayerColor_Blue)
	// e.Board.Cells[4][5] = b
	// e.Board.Cells[5][5] = &Ref{b}
	// for i := 1; i < 3; i++ {
	// 	e.Buildings = append(e.Buildings, &Building{Type: Government, Value: i})
	// }
	// for i := 0; i < 5; i++ {
	// 	e.Buildings = append(e.Buildings, NewBuildingHotel())
	// }
	// for i := 0; i < 5; i++ {
	//     e.Buildings = append(e.Buildings, &Building{Type: House})
	// }

	// for i := 0; i < 5; i++ {
	// 	e.Buildings = append(e.Buildings, NewBuildingShop())
	// }
}

var (
	totemPlaces = map[PlayerColor]map[BuildingType]Vector{
		PlayerColor_Red: {
			SmallTotem: Vector{3, 2},
			Totem:      Vector{4, 5},
		},
		PlayerColor_Green: {
			SmallTotem: Vector{3, 7},
			Totem:      Vector{4, 3},
		},
		PlayerColor_Blue: {
			SmallTotem: Vector{2, 2},
			Totem:      Vector{1, 5},
		},
		PlayerColor_Yellow: {
			SmallTotem: Vector{2, 7},
			Totem:      Vector{1, 3},
		},
	}
)
