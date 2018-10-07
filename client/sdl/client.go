package sdl

import (
	"runtime"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	windowWidth  = 1280
	windowHeight = 800

	windowXCenter = int32(windowWidth / 2)
	windowYCenter = int32(windowHeight / 2)

	fps          = 60
	deltaTimeFPS = time.Duration(time.Second / fps)

	resourcePath = "."
)

// Run - sdl game engine
func Run() error {
	runtime.LockOSThread()
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return errors.Wrap(err, "error on init sdl")
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Elasund", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, windowWidth, windowHeight, sdl.WINDOW_RESIZABLE)
	if err != nil {
		return errors.Wrap(err, "error on create window")
	}
	defer window.Destroy()

	ttf.Init()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return errors.Wrap(err, "error on create renderer")
	}
	defer renderer.Destroy()

	err = font.Initialize("Miramob.ttf")
	if err != nil {
		return errors.Wrap(err, "error on font initialize")
	}

	draw := &Draw{renderer: renderer}

	// var update Update

	var stages Stages
	stages.Add(StageNameMainMenu, NewMainMenuStage(&stages))
	stages.Add(StageNameIncome, DrawOnlyStage(func(draw *Draw) {
		draw.Clear(blue)
	}))
	stages.Add("3", DrawOnlyStage(func(draw *Draw) {
		draw.Clear(purple)
	}))

	// t := tm.TextureManager{}
	// err = t.Initialize(renderer)
	// if err != nil {
	// 	panic(err)
	// }
	// defer t.Close()

	// window.SetIcon(t.Icon)

	// count_players := 4
	// Elasund.Initialize(count_players)

	var event sdl.Event
	var isRunning = true
	var fpsLastTime time.Time

	var mousePoint = &MouseEvent{}
	// update.Mouse = &mousePoint
	// var mouse_over_map bool

	// var events = make(chan *Event, 1000)

	for _ = range time.Tick(deltaTimeFPS) {
		renderer.SetDrawColor(100, 10, 100, 255)
		renderer.Clear()

		stages.Draw(draw)

		// 	t.Common[tm.Common_Board].Draw(0, 0)

		ns := time.Since(fpsLastTime).Nanoseconds()
		currentFPS := int(int64(time.Second) / ns)
		fpsLastTime = time.Now()

		// draw_text("FPS: "+strconv.Itoa(int(currentFPS)), &point.Point{25, 5}, WHITE, 16)
		draw.Text("FPS: "+strconv.Itoa(currentFPS), Point{X: 25, Y: 5}, white, 16)

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				isRunning = false
			case *sdl.MouseButtonEvent:
				switch e.Button {
				case sdl.BUTTON_LEFT:
					switch e.State {
					case sdl.PRESSED:
						mousePoint.LMouseState = Pressed
						stages.Update(mousePoint)
					case sdl.RELEASED:
						mousePoint.LMouseState = Released
						stages.Update(mousePoint)
						mousePoint.LMouseState = NoneMouseState
					}
				}
			case *sdl.MouseMotionEvent:
				mousePoint.X = e.X
				mousePoint.Y = e.Y
				stages.Update(mousePoint)
			case *sdl.KeyboardEvent:
				switch e.Keysym.Sym {
				case sdl.K_ESCAPE, sdl.K_q:
					isRunning = false
				case sdl.K_1:
					stages.Next(StageNameMainMenu)
				case sdl.K_2:
					stages.Next("2")
				case sdl.K_3:
					stages.Next("3")
				}
			}
		} // events

		// 	draw_text(fmt.Sprintf("Mouse: (%d:%d)", mouse_point.X, mouse_point.Y), &point.Point{725, 20}, WHITE, 16)
		// 	if mouse_over_map {
		// 		mouse_cell_x, mouse_cell_y := get_cell(mouse_point)
		// 		draw_text(fmt.Sprintf("Cell: (%d:%d)", mouse_cell_x, mouse_cell_y), &point.Point{725, 40}, YELLOW, 16)
		// 	}
		// 	draw_text("FPS: "+strconv.Itoa(int(fps_now)), &point.Point{25, 5}, BLACK, 16)

		// 	for i := 0; i < 10; i++ {
		// 		for j := -1; j < 11; j++ {
		// 			draw_text(fmt.Sprintf("(%d, %d)", i, j), get_point(i, j).Move(8, 10), BLUE, 12)
		// 		}
		// 	}

		// 	t.Common[tm.Common_CornerTop].DrawPoint(get_point(count_players*2, -1))
		// 	t.Common[tm.Common_CornerBottom].DrawPoint(get_point(count_players*2, 9))

		// 	for _, b := range Elasund.Buildings {
		// 		if b.IsBuild {
		// 			t.Buildings[b.Type].DrawPoint(get_point(b.X, b.Y))
		// 		}
		// 	}

		// 	for i := 0; i < 9; i++ {
		// 		for j := 0; j < 10; j++ {
		// 			cell := Elasund.Board.Cells[i][j]
		// 			if cell != nil {
		// 				switch cell.GetType() {
		// 				case core.TileType_Building:
		// 					draw_rect(get_point(i, j), 1, 1, BLUE)
		// 				case core.TileType_Ref:
		// 					o := cell.(*core.Ref).Object
		// 					if o.GetType() == core.TileType_Building {
		// 						b := o.(*core.Building)
		// 						draw_rect(get_point(i, j), 1, 1, GREEN)
		// 						draw_line(
		// 							get_point(b.X, b.Y).Move(TILE_SIZE/2, TILE_SIZE/2),
		// 							get_point(i, j).Move(TILE_SIZE/2, TILE_SIZE/2),
		// 							GREEN,
		// 						)
		// 					}
		// 				}
		// 			}
		// 		}
		// 	}

		// 	if mouse_over_map {
		// 		t.Buildings[core.BuildingType_Fair].DrawPoint(get_point(get_cell(mouse_point)))
		// 	}

		if err = draw.Error(); err != nil {
			return err
		}

		renderer.Present()

		if !isRunning {
			break
		}
	}
	return nil
}
