package content

import "fmt"

const (
    C_ERR_UNKNOWN = "unknown"
    C_ERR_ACCOUNT_MISSING = "wrong account"
    C_ERR_PASSWORD_WRONG = "wrong password"
)

type DefinedError struct {
    Err string `json:"error"`
    Detail string `json:"detail,omitempty"`
}

func (self *DefinedError) Error() string {
    return fmt.Sprintf("Error: %s - %s", self.Err, self.Detail)
}

func BuildUnknownError(e error) *DefinedError {
    return &DefinedError{
        Err: C_ERR_UNKNOWN,
        Detail: e.Error(),
    }
}

func BuildWrappedError(code string, err error) *DefinedError {
    return &DefinedError{
        Err: code,
        Detail: err.Error(),
    }
}

func BuildError(code, detail string) *DefinedError {
    return &DefinedError{
        Err: code,
        Detail: detail,
    }
}

