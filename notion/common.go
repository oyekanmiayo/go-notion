package notion

// Type must be either "text", "mention" or "equation"
// Use Mention for Mention
type RichText struct {
	PlainText   string       `json:"plain_text,omitempty"`
	Href        string       `json:"href,omitempty"`
	Annotations *Annotations `json:"annotations,omitempty"`
	Type        string       `json:"type"`
	Text        Text         `json:"text,omitempty"`
	Mention     interface{}  `json:"mention,omitempty"`
	Equation    *Equation    `json:"equation,omitempty"`
}

// Color must be either "default", "gray", "brown", "orange", "yellow", "green", "blue", "purple", "pink", "red",
// "gray_background", "brown_background", "orange_background", "yellow_background", "green_background",
// "blue_background", "purple_background", "pink_background", or "red_background"
type Annotations struct {
	Bold          bool   `json:"bold,omitempty"`
	Italic        bool   `json:"italic,omitempty"`
	Strikethrough bool   `json:"strikethrough,omitempty"`
	Underline     bool   `json:"underline,omitempty"`
	Code          bool   `json:"code,omitempty"`
	Color         string `json:"color,omitempty"`
}

type Text struct {
	Content string `json:"content,omitempty"`
	Link    *Link  `json:"link,omitempty"`
}

// If, Link != nil, Type is always "url" and URL is always a web address.
// https://developers.notion.com/reference/rich-text#link-objects
type Link struct {
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}

// Use PageMention for PageMention
// Use DatabaseMention for DatabaseMention
// Use Date for Date
type Mention struct {
	Type            string      `json:"type"`
	UserMention     *User       `json:"user,omitempty"`
	PageMention     interface{} `json:"page,omitempty"`
	DatabaseMention interface{} `json:"database,omitempty"`
	Date            interface{} `json:"date,omitempty"`
}

type PageMention struct {
	ID string `json:"id,omitempty"`
}

type DatabaseMention struct {
	ID string `json:"id,omitempty"`
}

type Equation struct {
	Expression string `json:"expression,omitempty"`
}
