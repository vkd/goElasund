package texture_manager

import (
	"go_client/point"
	"path"

	"go_client/constants"
	"go_client/texture_manager/tm_errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var (
	_TEXTURES = [][]string{
		{"Board", "Board.png"},
		{"Corner_top", "Corner_top.png"},
		{"Corner_bottom", "Corner_bottom.png"},
	}
)

type TextureManager struct {
	is_initialized bool
	renderer       *sdl.Renderer

	textures map[string]*sdl.Texture
}

func (t *TextureManager) Initialize(renderer *sdl.Renderer) error {
	t.renderer = renderer
	t.textures = make(map[string]*sdl.Texture)

	var err error
	for _, texture_config := range _TEXTURES {
		err = t.load_file(texture_config[0], texture_config[1])
		if err != nil {
			return err
		}
	}

	t.is_initialized = true
	return nil
}

func (t *TextureManager) load_file(name string, path_file string) error {
	full_path := path.Join(constants.BASE_DIR, "textures", path_file)
	f := sdl.RWFromFile(full_path, "rb")
	defer f.RWclose()
	s, err := img.LoadPNG_RW(f)
	if err != nil {
		return &tm_errors.LoadTextureError{full_path, err}
	}
	defer s.Free()
	texture, err := t.renderer.CreateTextureFromSurface(s)
	if err != nil {
		return &tm_errors.CreateTextureError{name, err}
	}
	t.textures[name] = texture
	return nil
}

func (t *TextureManager) Get(name string) *sdl.Texture {
	return t.textures[name]
}

func (t *TextureManager) Draw(name string, x int32, y int32) {
	texture := t.Get(name)
	_, _, width, height, _ := texture.Query()
	t.renderer.Copy(texture, nil, &sdl.Rect{x, y, width, height})
}

func (t *TextureManager) DrawPoint(name string, p *point.Point) {
	t.Draw(name, int32(p.X), int32(p.Y))
}

func (t *TextureManager) Close() {
	for _, v := range t.textures {
		v.Destroy()
	}
}
