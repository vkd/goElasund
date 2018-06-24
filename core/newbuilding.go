package core

func NewBuildingChurch(value int) *Building {
	return &Building{
		Type:   Church,
		Width:  1,
		Height: 1,
		Value:  value,
	}
}

func NewBuildingDrawWell() *Building {
	return &Building{
		Type:   DrawWell,
		Width:  1,
		Height: 1,
	}
}

func NewBuildingFair() *Building {
	return &Building{
		Type:   Fair,
		Width:  2,
		Height: 1,
	}
}

func NewBuildingGovernment(value int) *Building {
	return &Building{
		Type:   Government,
		Width:  2,
		Height: 3,
		Value:  value,
	}
}
