package tm_errors

import "fmt"

type CreateTextureError struct {
	Name string
	Err  error
}

func (c *CreateTextureError) Error() string {
	return fmt.Sprintf("Error create texture \"%s\": %s", c.Name, c.Err)
}
