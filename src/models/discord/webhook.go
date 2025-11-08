package webhook

type WebhookResponse struct {
	MessageId string `json:"id,omitempty"`
	Message   string `json:"message,omitempty"`
}

type WebhookMessage struct {
	Username string  `json:"username,omitempty"`
	AvataUrl string  `json:"avatar_url,omitempty"`
	Content  string  `json:"content,omitempty"`
	Embeds   []Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Author      Author    `json:"author"`
	Title       string    `json:"title,omitempty"`
	Color       int       `json:"color,omitempty"`
	Description string    `json:"description,omitempty"`
	Fields      []Field   `json:"fields,omitempty"`
	Url         string    `json:"url,omitempty"`
	Image       SimpleUrl `json:"image"`
	Thumbnail   SimpleUrl `json:"thumbnail"`
	Footer      Footer    `json:"footer"`
	Timestamp   string    `json:"timestamp"`
}

type SimpleUrl struct {
	Url string `json:"url,omitempty"`
}

type Footer struct {
	Text    string `json:"text,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
}

type Field struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Author struct {
	Name    string `json:"name,omitempty"`
	IconUrl string `json:"icon_url,omitempty"`
	Url     string `json:"url,omitempty"`
}
