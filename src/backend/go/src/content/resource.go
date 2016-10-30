package content

import "fmt"

const (
	C_ERR_ARG_MISSING = "argument missing"
)

type ResourceError struct {
	Type string
	Detail error
}

func (self *ResourceError) Error() string {
	return fmt.Sprintf("Error(%s): %s", self.Type, self.Detail.Error())
}

type Resource interface {
	Validate() error
}
