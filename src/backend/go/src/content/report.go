package content


type decree struct {
	ID string `json:"id"`
	Content string `json:"content"`
	Source string `json:"source"`
	Link string `json:"link,omitempty"`
}

type event struct {
	ID string `json:"id"`
	Content string `json:"content"`
	Link string `json:"link,omitempty"`
}

type Report struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Conclusion string `json:"conclusion"`
	Decrees []*decree `json:"decrees,omitempty"`
	Cases []*event `json:"cases,omitempty"`
}

func (self *Report) Validate() error {
	return nil
}

