package sdl

import (
	"path"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	font Font
)

// Font of game with preload sizes
type Font struct {
	fonts map[int]*ttf.Font
}

// Initialize font with different sizes
func (f *Font) Initialize(filename string) error {
	f.fonts = make(map[int]*ttf.Font)

	var fnt *ttf.Font
	var err error
	for i := 6; i <= 24; i++ {
		fnt, err = ttf.OpenFont(path.Join(resourcePath, "fonts", filename), i)
		if err != nil {
			return errors.Wrap(err, "error on open font")
		}
		f.fonts[i] = fnt
	}

	return nil
}

// Size - return font with needed size
func (f *Font) Size(size int) *ttf.Font {
	return f.fonts[size]
}
