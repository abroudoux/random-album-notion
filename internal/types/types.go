package types

type NotionResponse struct {
	Object     string    `json:"object"`
	Results    []Block   `json:"results"`
	NextCursor *string   `json:"next_cursor"`
	HasMore    bool      `json:"has_more"`
}

type Block struct {
	Type string `json:"type"`
	ToDo *ToDo  `json:"to_do"`
}

type ToDo struct {
	Checked  bool       `json:"checked"`
	RichText []RichText `json:"rich_text"`
}

type RichText struct {
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}