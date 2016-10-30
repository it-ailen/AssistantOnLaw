package content

type Channel struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Deleted bool `json:"deleted"`
	CreatedTime uint64 `json:"created_time"`
}

func (self *Channel) Validate() error {

	return nil
}
