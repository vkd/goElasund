package go_client

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	font_helper "go_client/font"
	"go_client/texture_manager"

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

	tm := texture_manager.TextureManager{}
	err = tm.Initialize(renderer)
	if err != nil {
		panic(err)
	}

	font = new(font_helper.Font)
	err = font.Initialize("Miramob.ttf")
	if err != nil {
		panic(err)
	}

	var event sdl.Event
	var is_running = true
	var t time.Time

	var mouseX, mouseY int32
	var mouse_clicked bool

	for _ = range time.Tick(DELTA_FPS) {
		renderer.SetDrawColor(100, 10, 100, 255)
		renderer.Clear()

		tm.Draw("Board", 0, 0)

		ns := time.Since(t).Nanoseconds()
		fps_now := int64(time.Second) / ns
		t = time.Now()

		draw_text("FPS: "+strconv.Itoa(int(fps_now)), &Point{25, 5}, WHITE, 16)

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				is_running = false
			case *sdl.MouseButtonEvent:
				if t.Button == sdl.BUTTON_LEFT && t.State == sdl.PRESSED {
					mouse_clicked = !mouse_clicked
				}
			case *sdl.MouseMotionEvent:
				mouseX = t.X
				mouseY = t.Y
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					is_running = false
				}
				if t.Keysym.Sym == sdl.K_q {
					is_running = false
				}
			}
		}

		draw_text(fmt.Sprintf("Mouse: (%d:%d)", mouseX, mouseY), &Point{725, 20}, WHITE, 16)
		draw_text("FPS: "+strconv.Itoa(int(fps_now)), &Point{25, 5}, BLACK, 16)

		for i := 0; i < 10; i++ {
			for j := -1; j < 11; j++ {
				draw_text(fmt.Sprintf("(%d, %d)", i, j), get_point(i, j).Move(8, 10), BLUE, 12)
			}
		}

		if mouse_clicked {
			tm.Draw("Corner_top", mouseX, mouseY)
		}

		renderer.Present()

		if !is_running {
			break
		}
	}
	tm.Close()
	sdl.Quit()
}

func draw_text(text string, pos *Point, color sdl.Color, size int) {
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

func get_point(x int, y int) *Point {
	step := 50
	border := 1
	// X := 153 + x*(step+border)
	// Y := 184 + y*(step+border)
	return &Point{153 + x*(step+border), 184 + y*(step+border)}
}
