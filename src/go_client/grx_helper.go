package go_client

import (
	"go_client/point"

	"github.com/veandco/go-sdl2/sdl"
)

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
	return (mouse_point.X - 153) / TILE_STEP, (mouse_point.Y - 184) / TILE_STEP
}

func draw_rect(p *point.Point, width, height int, color sdl.Color) {
	W := int32(width*TILE_STEP - TILE_BORDER_SIZE)
	H := int32(height*TILE_STEP - TILE_BORDER_SIZE)

	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.FillRect(&sdl.Rect{
		X: int32(p.X),
		Y: int32(p.Y),
		W: W,
		H: LINE_WIDTH,
	})
	renderer.FillRect(&sdl.Rect{
		X: int32(p.X),
		Y: int32(p.Y) + H - LINE_WIDTH,
		W: W,
		H: LINE_WIDTH,
	})
	renderer.FillRect(&sdl.Rect{
		X: int32(p.X) + W - LINE_WIDTH,
		Y: int32(p.Y),
		W: LINE_WIDTH,
		H: H,
	})
	renderer.FillRect(&sdl.Rect{
		X: int32(p.X),
		Y: int32(p.Y),
		W: LINE_WIDTH,
		H: H,
	})
}

func draw_line(p1 *point.Point, p2 *point.Point, color sdl.Color) {
	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
}
