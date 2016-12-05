package definition

const (
	C_LAYOUT_SINGLE = "single"
	C_LAYOUT_MULTIPLE = "multiple"
)

type Entry struct {
	ID string `json:"id"`
	Text string `json:"text"`
	ChannelId string `json:"channel_id"`
	LayoutType string `json:"layout_type"`
}

func (self *Entry) Validate() error {

	return nil
}

