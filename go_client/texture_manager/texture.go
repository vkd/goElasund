package texture_manager

import (
	"goElasund/go_client/constants"
	"goElasund/go_client/point"
	"goElasund/go_client/texture_manager/tm_errors"
	"path"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Texture struct {
	Path string

	texture  *sdl.Texture
	renderer *sdl.Renderer
}

func (t *Texture) Init(rend *sdl.Renderer) error {
	t.renderer = rend
	var err error
	t.texture, err = t.load_file(t.Path)
	return err
}

func (t *Texture) Draw(x, y int) {
	_, _, width, height, _ := t.texture.Query()
	t.renderer.Copy(t.texture, nil, &sdl.Rect{int32(x), int32(y), width, height})
}

func (t *Texture) DrawPoint(p *point.Point) {
	t.Draw(p.X, p.Y)
}

func (t *Texture) Destroy() {
	t.texture.Destroy()
}

func (t *Texture) load_file(path_file string) (*sdl.Texture, error) {
	full_path := path.Join(constants.BASE_DIR, "textures", path_file)
	f := sdl.RWFromFile(full_path, "rb")
	defer f.RWclose()
	s, err := img.LoadPNGRW(f)
	if err != nil {
		return nil, &tm_errors.LoadTextureError{full_path, err}
	}
	defer s.Free()
	texture, err := t.renderer.CreateTextureFromSurface(s)
	if err != nil {
		return nil, &tm_errors.CreateTextureError{path_file, err}
	}
	return texture, nil
}
