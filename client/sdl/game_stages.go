package sdl

// MainMenuStage - stage with main menu of game
type MainMenuStage struct {
	buttons []*Button
}

// NewMainMenuStage - constructor
func NewMainMenuStage(sm StageManager) *MainMenuStage {
	return &MainMenuStage{
		buttons: []*Button{
			NewButton(Rect{X: 150, Y: 300, W: 260, H: 120}, "Start game", func() {
				sm.Next(StageNameIncome)
			}),
			NewButton(Rect{X: 450, Y: 300, W: 260, H: 120}, "Quit", func() {
				sm.Next(StageNameQuit)
			}),
		},
	}
}

func (m *MainMenuStage) Init() {}

// Update - implement updater
func (m *MainMenuStage) Update(e Eventer) {
	for i := range m.buttons {
		m.buttons[i].Update(e)
	}
}

// Draw main menu stage
func (m *MainMenuStage) Draw(draw *Draw) {
	draw.Clear(orange)
	draw.TextCenter("Main menu", Point{X: windowXCenter, Y: 40}, yellow, 24)

	for _, b := range m.buttons {
		b.Draw(draw)
	}
}

type DrawOnlyStage func(draw *Draw)

func (DrawOnlyStage) Update(e Eventer) {}

func (o DrawOnlyStage) Draw(draw *Draw) { o(draw) }
func (DrawOnlyStage) Init()             {}

type UpdateOnlyStage func(e Eventer)

func (u UpdateOnlyStage) Update(e Eventer) { u(e) }

func (UpdateOnlyStage) Draw(draw *Draw) {}
func (UpdateOnlyStage) Init()           {}

type InitOnlyStage func()

func (InitOnlyStage) Update(e Eventer) {}
func (InitOnlyStage) Draw(draw *Draw)  {}
func (i InitOnlyStage) Init()          { i() }
