package texture_manager

import (
	"path"

	"go_client/texture_manager/tm_errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

const (
	TEXTURES_PATH = "../../textures"
)

var (
	_TEXTURES = [][]string{
		{"Board", "Board.png"},
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
	full_path := path.Join(TEXTURES_PATH, path_file)
	f := sdl.RWFromFile(full_path, "rb")
	s, err := img.LoadPNG_RW(f)
	if err != nil {
		return &tm_errors.LoadTextureError{full_path, err}
	}
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
