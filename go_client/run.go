package go_client

import (
	"fmt"
	"goElasund/core"
	"strconv"
	"time"

	font_helper "goElasund/go_client/font"
	"goElasund/go_client/point"
	tm "goElasund/go_client/texture_manager"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Run - start game loop
func Run() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Elasund", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WINDOW_WIDTH, WINDOW_HEIGHT, sdl.WINDOW_RESIZABLE)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	ttf.Init()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	t := tm.TextureManager{}
	err = t.Initialize(renderer)
	if err != nil {
		panic(err)
	}

	window.SetIcon(t.Icon)

	font = new(font_helper.Font)
	err = font.Initialize("Miramob.ttf")
	if err != nil {
		panic(err)
	}

	count_players := 4
	Elasund.Initialize(core.AllPlayers)

	var event sdl.Event
	var is_running = true
	// var time_fps time.Time

	var mouse_point *point.Point = &point.Point{0, 0}
	var mouse_over_map bool

	var now time.Time
	// var prevNow time.Time = time.Now()
	// prevNow = time.Now().Add(-DELTA_FPS)
	// var workDelta, sleepDelta time.Duration
	var fps_now float64 = FPS

	var fpsTime time.Time = time.Now()
	// var fpsTimeNow time.Time

	tick := time.NewTicker(DELTA_FPS)
	for now = range tick.C {
		// now = time.Now()
		// log.Printf("    now: %v", now)
		// SLEEP
		// DELTA_FPS == workDelta + sleepDelta
		// workDelta = now.Sub(prevNow)

		fps_now = float64(time.Second) / float64(now.Sub(fpsTime).Nanoseconds())
		fpsTime = now

		// sleepDelta = DELTA_FPS - workDelta
		// log.Printf("\nworkDelta: %v\nsleepDelta: %v", workDelta, sleepDelta)
		// time.Sleep(sleepDelta)
		// prevNow = time.Now()
		// prevNow = now.Add(sleepDelta)

		// log.Printf("tick: %v", []map[string]interface{}{
		// 	{"fps_now": fps_now},
		// 	{"DELTA_FPS": DELTA_FPS},
		// 	{"sleepDelta": sleepDelta},
		// 	{"workDelta": workDelta},
		// 	{"work+sleep": workDelta + sleepDelta},
		// })

		// log.Printf("prevNow: %v (%v) (%v)", prevNow, int64(fps_now), fps_now)

		// WORK
		renderer.SetDrawColor(100, 10, 100, 255)
		renderer.Clear()

		t.Common[tm.Common_Board].Draw(0, 0)

		// fpsTimeNow = time.Now()
		draw_text("FPS: "+strconv.FormatInt(int64(fps_now+0.5), 10), &point.Point{25, 5}, WHITE, 16)

		// log.Printf("fps: %v (%v)", fpsTime, fps_now)

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				is_running = false
			case *sdl.MouseButtonEvent:
				if e.Button == sdl.BUTTON_LEFT && e.State == sdl.PRESSED {
					// cell_x, cell_y := get_cell(mouse_point)
					// if mouse_over_map {
					// if Elasund.CheckBuild(core.Hotel, cell_x, cell_y, core.PlayerColor_Blue) {
					// 	Elasund.Build(core.Hotel, cell_x, cell_y, core.PlayerColor_Blue)
					// }
					// }
				}
				if e.Button == sdl.BUTTON_RIGHT && e.State == sdl.PRESSED {
					// cell_x, cell_y := get_cell(mouse_point)
					// if mouse_over_map {
					// if Elasund.CheckBuild(core.Fair, cell_x, cell_y, core.PlayerColor_Blue) {
					// 	Elasund.Build(core.Fair, cell_x, cell_y, core.PlayerColor_Blue)
					// }
					// }
				}
			case *sdl.MouseMotionEvent:
				mouse_point = &point.Point{int(e.X), int(e.Y)}

				mouse_over_map = false
				if mouse_point.X >= 153 && mouse_point.Y >= 184 {
					if mouse_point.X < 612 && mouse_point.Y < 694 {
						mouse_over_map = true
					}
				}
			case *sdl.KeyboardEvent:
				if e.Keysym.Sym == sdl.K_ESCAPE {
					is_running = false
				}
				if e.Keysym.Sym == sdl.K_q {
					is_running = false
				}
				if e.Keysym.Sym == sdl.K_r {
					for i, _ := range Elasund.Buildings {
						Elasund.Buildings[i].IsBuilded = false
					}
				}
			}
		}

		drawingArea := &Region{}
		boardArea := NewRegion(drawingArea, 0, 0)
		drawingArea.Add(boardArea)

		drawingArea.Draw(renderer)

		draw_text(fmt.Sprintf("Mouse: (%d:%d)", mouse_point.X, mouse_point.Y), &point.Point{725, 20}, WHITE, 16)
		if mouse_over_map {
			cell := get_cell(mouse_point)
			mouse_cell_x, mouse_cell_y := cell.X, cell.Y
			draw_text(fmt.Sprintf("Cell: (%d:%d)", mouse_cell_x, mouse_cell_y), &point.Point{725, 40}, YELLOW, 16)
		}
		draw_text("FPS: "+strconv.Itoa(int(fps_now)), &point.Point{25, 5}, BLACK, 16)

		for i := 0; i < 10; i++ {
			for j := -1; j < 11; j++ {
				draw_text(fmt.Sprintf("(%d, %d)", i, j), get_point(point.Point{i, j}).Move(8, 10), BLUE, 12)
			}
		}

		t.Common[tm.Common_CornerTop].DrawPoint(get_point(point.Point{count_players * 2, -1}))
		t.Common[tm.Common_CornerBottom].DrawPoint(get_point(point.Point{count_players * 2, 9}))

		for _, b := range Elasund.Buildings {
			if b.IsBuilded {
				// t.Buildings[b.Type].DrawPoint(get_point(b.X, b.Y))
			}
		}

		for _, b := range Elasund.Buildings {
			if !b.IsBuilded {
				continue
			}
			t.Buildings[b.Type][b.Index].DrawPoint(get_point(point.Point(b.Vector)))
		}

		if mouse_over_map {
			t.Buildings[core.DrawWell][0].DrawPoint(get_point(get_cell(mouse_point)))
		}

		renderer.Present()

		if !is_running {
			break
		}
	}
	t.Close()
	sdl.Quit()
}
