package font

import (
	"goElasund/go_client/constants"
	"path"

	"github.com/veandco/go-sdl2/ttf"
)

type Font struct {
	fonts map[int]*ttf.Font
}

func (f *Font) Initialize(font_path string) error {
	f.fonts = make(map[int]*ttf.Font)

	var fnt *ttf.Font
	var err error
	for i := 6; i <= 20; i++ {
		fnt, err = ttf.OpenFont(path.Join(constants.BASE_DIR, "fonts", font_path), i)
		if err != nil {
			return err
		}
		f.fonts[i] = fnt
	}

	return nil
}

func (f *Font) Size(size int) *ttf.Font {
	return f.fonts[size]
}
