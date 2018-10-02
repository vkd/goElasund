package sdl

// MainMenuGameStage - main menu stage func
func MainMenuGameStage(draw *Draw) Stage {
	return func() {
		draw.Clear(orange)
		draw.Text("Main menu", 25, 35, yellow, 24)

		// draw.Button()
	}
}
