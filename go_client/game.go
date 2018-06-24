package go_client

import (
	"goElasund/core"
	"time"

	font_helper "goElasund/go_client/font"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	FPS = 60

	WINDOW_WIDTH  = 1280
	WINDOW_HEIGHT = 800
)

var (
	DELTA_FPS = time.Duration(int64(time.Second) / FPS)

	font     *font_helper.Font
	renderer *sdl.Renderer

	Elasund = core.NewElasund()

	BLACK  = sdl.Color{0, 0, 0, 255}
	WHITE  = sdl.Color{255, 255, 255, 255}
	PURPLE = sdl.Color{100, 10, 100, 255}
	BLUE   = sdl.Color{10, 10, 255, 255}
	YELLOW = sdl.Color{255, 255, 100, 255}
	GREEN  = sdl.Color{10, 255, 10, 255}

	TILE_SIZE        = 50
	TILE_BORDER_SIZE = 1
	TILE_STEP        = TILE_SIZE + TILE_BORDER_SIZE

	TILE_START_X = 153
	TILE_START_Y = 184

	LINE_WIDTH int32 = 3
)

// type Game struct {
// 	is_running bool
// }

// func (g *Game) Init() {

// }

// func (g *Game) MainLoop() {
// 	for _ = range time.Tick(DELTA_FPS) {

// 		g.update_events()

// 		if !g.is_running {
// 			break
// 		}
// 	}
// }

// func (g *Game) update_events() {
// 	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
// 		switch e := event.(type) {
// 		case *sdl.QuitEvent:
// 			g.stop_loop()
// 		case *sdl.MouseButtonEvent:
// 			if e.Button == sdl.BUTTON_LEFT && e.State == sdl.PRESSED {
// 				cell_x, cell_y := get_cell(mouse_point)
// 				if mouse_over_map {
// 					if Elasund.CheckBuild(core.Hotel, cell_x, cell_y, core.Blue) {
// 						Elasund.Build(core.Hotel, cell_x, cell_y, core.Blue)
// 					}
// 				}
// 			}
// 			if e.Button == sdl.BUTTON_RIGHT && e.State == sdl.PRESSED {
// 				cell_x, cell_y := get_cell(mouse_point)
// 				if mouse_over_map {
// 					if Elasund.CheckBuild(core.Fair, cell_x, cell_y, core.Blue) {
// 						Elasund.Build(core.Fair, cell_x, cell_y, core.Blue)
// 					}
// 				}
// 			}
// 		case *sdl.MouseMotionEvent:
// 			mouse_point = &point.Point{int(e.X), int(e.Y)}

// 			mouse_over_map = false
// 			if mouse_point.X >= 153 && mouse_point.Y >= 184 {
// 				if mouse_point.X < 612 && mouse_point.Y < 694 {
// 					mouse_over_map = true
// 				}
// 			}
// 		case *sdl.KeyDownEvent:
// 			if e.Keysym.Sym == sdl.K_ESCAPE {
// 				g.stop_loop()
// 			}
// 			if e.Keysym.Sym == sdl.K_q {
// 				g.stop_loop()
// 			}
// 			if e.Keysym.Sym == sdl.K_r {
// 				for i, _ := range Elasund.Buildings {
// 					Elasund.Buildings[i].IsBuild = false
// 				}
// 			}
// 		}
// 	}
// }

// func (g *Game) stop_loop() {
// 	g.is_running = false
// }
