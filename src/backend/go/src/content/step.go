package content

const (
	C_ST_option = "option"
	C_ST_report = "report"
)

type Option struct {
	ID string `json:"id"`
	Text string `json:"text"`
	ParentId string `json:"parent_id"`
	Type string `json:"type"`
	Report *Report `json:"report,omitempty"`
}

func (self *Option) Validate() error {
	return nil
}

