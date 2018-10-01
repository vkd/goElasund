package texture_manager

import (
	"fmt"
	"goElasund/core"
	"path"

	"goElasund/go_client/constants"
	"goElasund/go_client/texture_manager/tm_errors"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
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
	}

	building_textures = []*struct {
		T core.BuildingType
		P string
	}{
		{core.DrawWell, "Buildings/DrawWell.png"},
		{core.Fair, "Buildings/Fair.png"},
		{core.Hotel, "Buildings/Hotel.png"},
		{core.Shop, "Buildings/Shop.png"},
	}

	usersTextures = []*struct {
		T core.BuildingType
		P string
	}{
		{core.House, "Buildings/House%d.png"},
		{core.SmallTotem, "Buildings/SmallTotem%d.png"},
		{core.Totem, "Buildings/Totem%d.png"},
		{core.Workshop, "Buildings/WorkShop%d.png"},
	}
)

type TextureManager struct {
	is_initialized bool
	renderer       *sdl.Renderer

	Common    map[Common_TextureType]*Texture
	Buildings map[core.BuildingType][]*Texture

	Icon *sdl.Surface
}

func (t *TextureManager) Initialize(renderer *sdl.Renderer) error {
	t.renderer = renderer
	t.Common = make(map[Common_TextureType]*Texture)
	t.Buildings = make(map[core.BuildingType][]*Texture)

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

	for _, v := range building_textures {
		t.Buildings[v.T] = []*Texture{&Texture{Path: v.P}}
	}
	for _, v := range usersTextures {
		for _, p := range core.AllPlayers {
			t.Buildings[v.T] = append(t.Buildings[v.T], &Texture{Path: fmt.Sprintf(v.P, int(p))})
		}
	}

	t.Buildings[core.Government] = append(t.Buildings[core.Government], nil)
	for i := 0; i < 3; i++ {
		t.Buildings[core.Government] = append(t.Buildings[core.Government], &Texture{
			Path: fmt.Sprintf("Buildings/Government%d.png", i+1),
		})
	}

	t.Buildings[core.Church] = append(t.Buildings[core.Church], nil)
	for i := 0; i < 9; i++ {
		t.Buildings[core.Church] = append(t.Buildings[core.Church], &Texture{
			Path: fmt.Sprintf("Buildings/Church%d.png", i+1),
		})
	}

	for _, tt := range t.Buildings {
		for _, tex := range tt {
			err = tex.Init(renderer)
			if err != nil {
				return err
			}
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
	s, err := img.LoadPNGRW(f)
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
