package sdl

import (
	"fmt"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

// Draw primitives
type Draw struct {
	renderer *sdl.Renderer
	errsmsg  []string
}

// Errors - return errors from drawer
func (d *Draw) Error() error {
	if len(d.errsmsg) == 0 {
		return nil
	}
	return fmt.Errorf("errors on draw: [%s]", strings.Join(d.errsmsg, ", "))
}

func (d *Draw) err(err error, msg string) {
	if err != nil {
		d.errsmsg = append(d.errsmsg, msg)
	}
}

// Text - draw text on sdl.Renderer
func (d *Draw) Text(text string, x, y int32, color sdl.Color, size int) {
	surf, err := font.Size(size).RenderUTF8Blended(text, color)
	if err != nil {
		d.err(err, "error on render text")
		return
	}
	defer surf.Free()
	textTexture, err := d.renderer.CreateTextureFromSurface(surf)
	if err != nil {
		d.err(err, "error on create texture")
		return
	}
	defer textTexture.Destroy()
	err = d.renderer.Copy(textTexture, nil, &sdl.Rect{X: x, Y: y, W: surf.W, H: surf.H})
	if err != nil {
		d.err(err, "error on copy texture to surface")
		return
	}
}

// Clear renderer surface
func (d *Draw) Clear(background sdl.Color) {
	err := d.renderer.SetDrawColor(
		background.R,
		background.G,
		background.B,
		background.A,
	)
	d.err(err, "error on set draw color")

	err = d.renderer.Clear()
	d.err(err, "error on clear renderer")
}

// // TileRect - draw rectangle on
// func (d *Draw) TileRect(x, y int32, width, height int, color sdl.Color) {
// 	W := int32(width*TILE_STEP - TILE_BORDER_SIZE)
// 	H := int32(height*TILE_STEP - TILE_BORDER_SIZE)

// 	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
// 	renderer.FillRect(&sdl.Rect{
// 		X: x,
// 		Y: y,
// 		W: W,
// 		H: LINE_WIDTH,
// 	})
// 	renderer.FillRect(&sdl.Rect{
// 		X: x,
// 		Y: y + H - LINE_WIDTH,
// 		W: W,
// 		H: LINE_WIDTH,
// 	})
// 	renderer.FillRect(&sdl.Rect{
// 		X: x + W - LINE_WIDTH,
// 		Y: y,
// 		W: LINE_WIDTH,
// 		H: H,
// 	})
// 	renderer.FillRect(&sdl.Rect{
// 		X: x,
// 		Y: y,
// 		W: LINE_WIDTH,
// 		H: H,
// 	})
// }

// func draw_line(p1 *point.Point, p2 *point.Point, color sdl.Color) {
// 	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
// 	renderer.DrawLine(p1.X, p1.Y, p2.X, p2.Y)
// }
