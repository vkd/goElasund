package go_client

import (
	"core"
	"fmt"
	"runtime"
	"strconv"
	"time"

	font_helper "go_client/font"
	"go_client/point"
	tm "go_client/texture_manager"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func Run() {
	runtime.LockOSThread()
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
	Elasund.Initialize(count_players)

	var event sdl.Event
	var is_running = true
	var time_fps time.Time

	var mouse_point *point.Point = &point.Point{0, 0}
	var mouse_over_map bool

	for _ = range time.Tick(DELTA_FPS) {
		renderer.SetDrawColor(100, 10, 100, 255)
		renderer.Clear()

		t.Common[tm.Common_Board].Draw(0, 0)

		ns := time.Since(time_fps).Nanoseconds()
		fps_now := int64(time.Second) / ns
		time_fps = time.Now()

		draw_text("FPS: "+strconv.Itoa(int(fps_now)), &point.Point{25, 5}, WHITE, 16)

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				is_running = false
			case *sdl.MouseButtonEvent:
				if e.Button == sdl.BUTTON_LEFT && e.State == sdl.PRESSED {
					cell_x, cell_y := get_cell(mouse_point)
					if mouse_over_map {
						if Elasund.CheckBuild(core.BuildingType_Hotel, cell_x, cell_y, core.PlayerColor_Blue) {
							Elasund.Build(core.BuildingType_Hotel, cell_x, cell_y, core.PlayerColor_Blue)
						}
					}
				}
				if e.Button == sdl.BUTTON_RIGHT && e.State == sdl.PRESSED {
					cell_x, cell_y := get_cell(mouse_point)
					if mouse_over_map {
						if Elasund.CheckBuild(core.BuildingType_Fair, cell_x, cell_y, core.PlayerColor_Blue) {
							Elasund.Build(core.BuildingType_Fair, cell_x, cell_y, core.PlayerColor_Blue)
						}
					}
				}
			case *sdl.MouseMotionEvent:
				mouse_point = &point.Point{int(e.X), int(e.Y)}

				mouse_over_map = false
				if mouse_point.X >= 153 && mouse_point.Y >= 184 {
					if mouse_point.X < 612 && mouse_point.Y < 694 {
						mouse_over_map = true
					}
				}
			case *sdl.KeyDownEvent:
				if e.Keysym.Sym == sdl.K_ESCAPE {
					is_running = false
				}
				if e.Keysym.Sym == sdl.K_q {
					is_running = false
				}
				if e.Keysym.Sym == sdl.K_r {
					for i, _ := range Elasund.Buildings {
						Elasund.Buildings[i].IsBuild = false
					}
				}
			}
		}

		draw_text(fmt.Sprintf("Mouse: (%d:%d)", mouse_point.X, mouse_point.Y), &point.Point{725, 20}, WHITE, 16)
		if mouse_over_map {
			mouse_cell_x, mouse_cell_y := get_cell(mouse_point)
			draw_text(fmt.Sprintf("Cell: (%d:%d)", mouse_cell_x, mouse_cell_y), &point.Point{725, 40}, YELLOW, 16)
		}
		draw_text("FPS: "+strconv.Itoa(int(fps_now)), &point.Point{25, 5}, BLACK, 16)

		for i := 0; i < 10; i++ {
			for j := -1; j < 11; j++ {
				draw_text(fmt.Sprintf("(%d, %d)", i, j), get_point(i, j).Move(8, 10), BLUE, 12)
			}
		}

		t.Common[tm.Common_CornerTop].DrawPoint(get_point(count_players*2, -1))
		t.Common[tm.Common_CornerBottom].DrawPoint(get_point(count_players*2, 9))

		for _, b := range Elasund.Buildings {
			if b.IsBuild {
				t.Buildings[b.Type].DrawPoint(get_point(b.X, b.Y))
			}
		}

		for i := 0; i < 9; i++ {
			for j := 0; j < 10; j++ {
				cell := Elasund.Board.Cells[i][j]
				if cell != nil {
					switch cell.GetType() {
					case core.TileType_Building:
						draw_rect(get_point(i, j), 1, 1, BLUE)
					case core.TileType_Ref:
						o := cell.(*core.Ref).Object
						if o.GetType() == core.TileType_Building {
							b := o.(*core.Building)
							draw_rect(get_point(i, j), 1, 1, GREEN)
							draw_line(
								get_point(b.X, b.Y).Move(TILE_SIZE/2, TILE_SIZE/2),
								get_point(i, j).Move(TILE_SIZE/2, TILE_SIZE/2),
								GREEN,
							)
						}
					}
				}
			}
		}

		if mouse_over_map {
			t.Buildings[core.BuildingType_Fair].DrawPoint(get_point(get_cell(mouse_point)))
		}

		renderer.Present()

		if !is_running {
			break
		}
	}
	t.Close()
	sdl.Quit()
}
