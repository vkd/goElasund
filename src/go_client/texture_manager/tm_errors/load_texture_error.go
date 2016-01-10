package tm_errors

import (
	"fmt"
)

type LoadTextureError struct {
	Path string
	Err  error
}

func (l *LoadTextureError) Error() string {
	return fmt.Sprintf("Error load file \"%s\": %s", l.Path, l.Err)
}
