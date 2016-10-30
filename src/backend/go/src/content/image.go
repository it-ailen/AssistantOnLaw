package content

import "fmt"

type Image struct {
	URI string `json:"uri"`
}

func (self *Image) Validate() error {
	if len(self.URI) == 0 {
		return &ResourceError{
			Type: C_ERR_ARG_MISSING,
			Detail: fmt.Errorf("Missing `%s`", "URI"),
		}
	}
	return nil
}
