package core

import (
	"math/rand"
	"time"
)

type Elasund struct {
	random *rand.Rand

	Buildings []*Building
}

func NewElasund() (e *Elasund) {
	e = new(Elasund)
	e.random = rand.New(rand.NewSource(time.Now().UnixNano()))

	return e
}

func (e *Elasund) Build(building_type BuildingType, x, y int, color PlayerColor) {
	for _, b := range e.Buildings {
		if b.OnMap {
			continue
		}
		if b.Type != building_type {
			continue
		}
		b.Build(x, y, color)
		return
	}
}

func (e *Elasund) Initialize() {
	// for i := 1; i <= 9; i++ {
	// 	e.Buildings = append(e.Buildings, &Building{Type: Church, Value: i})
	// }
	for i := 0; i < 4; i++ {
		e.Buildings = append(e.Buildings, &Building{Type: DrawWell})
	}
	for i := 0; i < 4; i++ {
		e.Buildings = append(e.Buildings, &Building{Type: Fair})
	}
	// for i := 1; i < 3; i++ {
	// 	e.Buildings = append(e.Buildings, &Building{Type: Government, Value: i})
	// }
	for i := 0; i < 5; i++ {
		e.Buildings = append(e.Buildings, &Building{Type: Hotel})
	}
	// for i := 0; i < 5; i++ {
	//     e.Buildings = append(e.Buildings, &Building{Type: House})
	// }

	for i := 0; i < 5; i++ {
		e.Buildings = append(e.Buildings, &Building{Type: Shop})
	}
}
