package content

type Issue struct {
    ID string `json:"id"`
    CreatedTime int64 `json:"created_time"`
	Client struct {
		Name    string `json:"name"`
		Contact string `json:"contact"`
	} `json:"client"`
	Detail struct {
		Desc        string   `json:"description"`
		Attachments []string `json:"attachments"`
	} `json:"detail"`
    Status string `json:"status"`
    Solution string `json:"solution"`
    Tags []string `json:"tags"`
}

func (self *Issue) Validate() error {

	return nil
}
