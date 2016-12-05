package definition

const (
    C_ACC_TYPE_CUSTOMER = "customer"
    C_ACC_TYPE_SUPER = "super"
)

type Account struct {
    ID string `json:"id"`
    Account string `json:"account"`
    Password string `json:"-"`
    Type string `json:"type"`
    Nick string `json:"nick,omitempty"`
    Contact string `json:"contact,omitempty"`
    Etc string `json:"etc"`
    Session string `json:"-"`
}
