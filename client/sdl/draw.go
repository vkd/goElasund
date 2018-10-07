package sdl

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

// Draw primitives
type Draw struct {
	renderer *sdl.Renderer
	errsmsg  []string
}

// Drawer interface for object what can be drawed on screen
type Drawer interface {
	Draw(draw *Draw)
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
		d.errsmsg = append(d.errsmsg, errors.Wrap(err, msg).Error())
	}
}

// Text - draw text on sdl.Renderer
func (d *Draw) Text(text string, p Point, color sdl.Color, size int) {
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
	err = d.renderer.Copy(textTexture, nil, &sdl.Rect{X: p.X, Y: p.Y, W: surf.W, H: surf.H})
	if err != nil {
		d.err(err, "error on copy texture to surface")
		return
	}
}

// TextCenter - draw text on sdl.Renderer
func (d *Draw) TextCenter(text string, p Point, color sdl.Color, size int) {
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
	err = d.renderer.Copy(textTexture, nil, &sdl.Rect{X: p.X - int32(surf.W/2), Y: p.Y - int32(surf.H/2), W: surf.W, H: surf.H})
	if err != nil {
		d.err(err, "error on copy texture to surface")
		return
	}
}

// Clear renderer surface
func (d *Draw) Clear(background sdl.Color) {
	d.setColor(background)

	err := d.renderer.Clear()
	d.err(err, "error on clear renderer")
}

// RectByBorder - draw rectangle
func (d *Draw) RectByBorder(r Rect, thickness int32, color sdl.Color) {
	d.setColor(color)

	d.renderer.FillRect(Rect{X: r.X, Y: r.Y, W: r.W, H: thickness}.SDLRect())
	d.renderer.FillRect(Rect{X: r.X, Y: r.Bottom() - thickness, W: r.W, H: thickness}.SDLRect())

	d.renderer.FillRect(Rect{X: r.X, Y: r.Y, W: thickness, H: r.H}.SDLRect())
	d.renderer.FillRect(Rect{X: r.Right() - thickness, Y: r.Y, W: thickness, H: r.H}.SDLRect())
}

// RectByFill - draw rectangle by fill two colored rectangle
func (d *Draw) RectByFill(r Rect, thickness int32, colorBorder, colorInner sdl.Color) {
	d.setColor(colorBorder)
	d.renderer.FillRect(r.SDLRect())

	d.setColor(colorInner)
	d.renderer.FillRect(&sdl.Rect{
		X: r.X + thickness,
		Y: r.Y + thickness,
		W: r.W - 2*thickness,
		H: r.H - 2*thickness,
	})
}

func (d *Draw) setColor(color sdl.Color) {
	err := d.renderer.SetDrawColor(
		color.R,
		color.G,
		color.B,
		color.A,
	)
	d.err(err, "error on set draw color")
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
