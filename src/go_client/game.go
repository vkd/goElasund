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

	Elasund.Initialize()

	var event sdl.Event
	var is_running = true
	var time_fps time.Time

	var mouse_point *point.Point = &point.Point{0, 0}
	var mouse_clicked bool

	count_players := 4

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
					// mouse_clicked = !mouse_clicked
					cell_x, cell_y := get_cell(mouse_point)
					Elasund.Build(core.Hotel, cell_x, cell_y, core.Blue)
				}
			case *sdl.MouseMotionEvent:
				mouse_point = &point.Point{int(e.X), int(e.Y)}
			case *sdl.KeyDownEvent:
				if e.Keysym.Sym == sdl.K_ESCAPE {
					is_running = false
				}
				if e.Keysym.Sym == sdl.K_q {
					is_running = false
				}
			}
		}

		draw_text(fmt.Sprintf("Mouse: (%d:%d)", mouse_point.X, mouse_point.Y), &point.Point{725, 20}, WHITE, 16)
		mouse_cell_x, mouse_cell_y := get_cell(mouse_point)
		draw_text(fmt.Sprintf("Cell: (%d:%d)", mouse_cell_x, mouse_cell_y), &point.Point{725, 40}, WHITE, 16)
		draw_text("FPS: "+strconv.Itoa(int(fps_now)), &point.Point{25, 5}, BLACK, 16)

		for i := 0; i < 10; i++ {
			for j := -1; j < 11; j++ {
				draw_text(fmt.Sprintf("(%d, %d)", i, j), get_point(i, j).Move(8, 10), BLUE, 12)
			}
		}

		t.Common[tm.Common_CornerTop].DrawPoint(get_point(count_players*2, -1))
		t.Common[tm.Common_CornerBottom].DrawPoint(get_point(count_players*2, 9))

		for _, b := range Elasund.Buildings {
			if b.OnMap {
				t.Common[tm.Common_Hotel].DrawPoint(get_point(b.X, b.Y))
			}
		}

		if mouse_point.X >= 153 && mouse_point.Y >= 184 {
			if mouse_point.X < 612 && mouse_point.Y < 694 {
				t.Common[tm.Common_Hotel].DrawPoint(get_point(get_cell(mouse_point)))
			}
		}

		if mouse_clicked {
			t.Common[tm.Common_Hotel].DrawPoint(mouse_point)
		}

		renderer.Present()

		if !is_running {
			break
		}
	}
	t.Close()
	sdl.Quit()
}

func draw_text(text string, pos *point.Point, color sdl.Color, size int) {
	surf, err := font.Size(size).RenderUTF8_Blended(text, color)
	if err != nil {
		panic(err)
	}
	defer surf.Free()
	text_texture, err := renderer.CreateTextureFromSurface(surf)
	if err != nil {
		panic(err)
	}
	defer text_texture.Destroy()
	err = renderer.Copy(text_texture, nil, &sdl.Rect{int32(pos.X), int32(pos.Y), surf.W, surf.H})
	if err != nil {
		panic(err)
	}
}

func get_point(x int, y int) *point.Point {
	step := 50
	border := 1
	// X := 153 + x*(step+border)
	// Y := 184 + y*(step+border)
	return &point.Point{153 + x*(step+border), 184 + y*(step+border)}
}

func get_cell(mouse_point *point.Point) (int, int) {
	return (mouse_point.X - 153) / 51, (mouse_point.Y - 184) / 51
}
