package main

import (
	"runtime"
	"strconv"
	"time"

	"go_client/texture_manager"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

const (
	FPS = 30

	WINDOW_WIDTH  = 1280
	WINDOW_HEIGHT = 800
)

var (
	DELTA_FPS = time.Duration(int64(time.Second) / FPS)

	font     *ttf.Font
	renderer *sdl.Renderer

	BLACK  = sdl.Color{0, 0, 0, 255}
	WHITE  = sdl.Color{255, 255, 255, 255}
	PURPLE = sdl.Color{100, 10, 100, 255}
)

func main() {
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

	font, err = ttf.OpenFont("../../fonts/Miramob.ttf", 16)
	if err != nil {
		panic(err)
	}

	var event sdl.Event
	var is_running = true
	var t time.Time

	for _ = range time.Tick(DELTA_FPS) {
		renderer.SetDrawColor(100, 10, 100, 255)
		renderer.Clear()

		tm.Draw("Board", 0, 0)

		ns := time.Since(t).Nanoseconds()
		fps_now := int64(time.Second) / ns
		t = time.Now()

		draw_text("FPS: "+strconv.Itoa(int(fps_now)), 25, 5, WHITE)

		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				is_running = false
			case *sdl.MouseButtonEvent:
				if t.Button == sdl.BUTTON_LEFT && t.State == sdl.PRESSED {
					mouse_rect := sdl.Rect{t.X, t.Y, 5, 5}
					var err error
					err = renderer.SetDrawColor(10, 255, 10, 255)
					if err != nil {
						panic(err)
					}
					err = renderer.FillRect(&mouse_rect)
					if err != nil {
						panic(err)
					}
				}
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					is_running = false
				}
				if t.Keysym.Sym == sdl.K_q {
					is_running = false
				}
			}
		}
		renderer.Present()

		if !is_running {
			break
		}
	}
	sdl.Quit()
}

func draw_text(text string, x int32, y int32, color sdl.Color) {
	surf, err := font.RenderUTF8_Blended(text, color)
	if err != nil {
		panic(err)
	}
	text_texture, err := renderer.CreateTextureFromSurface(surf)
	if err != nil {
		panic(err)
	}
	err = renderer.Copy(text_texture, nil, &sdl.Rect{x, y, surf.W, surf.H})
	if err != nil {
		panic(err)
	}
}
