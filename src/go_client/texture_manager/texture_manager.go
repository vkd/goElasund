package texture_manager

import (
	"path"

	"go_client/constants"
	"go_client/texture_manager/tm_errors"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

type TextureManager struct {
	is_initialized bool
	renderer       *sdl.Renderer

	Common map[int]*Texture
}

func (t *TextureManager) Initialize(renderer *sdl.Renderer) error {
	t.renderer = renderer
	t.Common = make(map[int]*Texture)

	t.Common[Common_Board] = &Texture{Path: "Board.png"}
	t.Common[Common_ChurchBack] = &Texture{Path: "Church0.png"}
	t.Common[Common_CornerBottom] = &Texture{Path: "Corner_bottom.png"}
	t.Common[Common_CornerTop] = &Texture{Path: "Corner_top.png"}

	var tex *Texture
	var err error
	for _, tex = range t.Common {
		err = tex.Init(renderer)
		if err != nil {
			return err
		}
	}

	t.is_initialized = true
	return nil
}

func (t *TextureManager) load_file(path_file string) (*sdl.Texture, error) {
	full_path := path.Join(constants.BASE_DIR, "textures", path_file)
	f := sdl.RWFromFile(full_path, "rb")
	defer f.RWclose()
	s, err := img.LoadPNG_RW(f)
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

func (t *TextureManager) Close() {
	for _, v := range t.Common {
		v.Destroy()
	}
}
