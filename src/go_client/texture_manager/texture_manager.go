package texture_manager

import (
	"path"

	"go_client/constants"
	"go_client/texture_manager/tm_errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

var (
	common_textures = []*struct {
		T Common_TextureType
		S string
	}{
		{Common_Board, "Board.png"},
		{Common_ChurchBack, "Church0.png"},
		{Common_CornerBottom, "Corner_bottom.png"},
		{Common_CornerTop, "Corner_top.png"},

		{Common_Hotel, "Buildings/Hotel.png"},
	}
)

type TextureManager struct {
	is_initialized bool
	renderer       *sdl.Renderer

	Common map[Common_TextureType]*Texture

	Icon *sdl.Surface
}

func (t *TextureManager) Initialize(renderer *sdl.Renderer) error {
	t.renderer = renderer
	t.Common = make(map[Common_TextureType]*Texture)

	for _, v := range common_textures {
		t.Common[v.T] = &Texture{Path: v.S}
	}

	var tex *Texture
	var err error
	for _, tex = range t.Common {
		err = tex.Init(renderer)
		if err != nil {
			return err
		}
	}

	t.Icon, err = t.load_file_as_surface("icon.png")
	if err != nil {
		return err
	}

	t.is_initialized = true
	return nil
}

func (t *TextureManager) load_file(path_file string) (*sdl.Texture, error) {
	s, err := t.load_file_as_surface(path_file)
	if err != nil {
		return nil, &tm_errors.LoadTextureError{path_file, err}
	}
	defer s.Free()
	texture, err := t.renderer.CreateTextureFromSurface(s)
	if err != nil {
		return nil, &tm_errors.CreateTextureError{path_file, err}
	}
	return texture, nil
}

func (t *TextureManager) load_file_as_surface(path_file string) (*sdl.Surface, error) {
	full_path := path.Join(constants.BASE_DIR, "textures", path_file)
	f := sdl.RWFromFile(full_path, "rb")
	defer f.RWclose()
	s, err := img.LoadPNG_RW(f)
	if err != nil {
		return nil, &tm_errors.LoadTextureError{full_path, err}
	}
	return s, nil
}

func (t *TextureManager) Close() {
	for _, v := range t.Common {
		v.Destroy()
	}
	t.Icon.Free()
}
