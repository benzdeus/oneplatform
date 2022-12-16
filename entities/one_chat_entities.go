package one_entities

type OneChatService interface {
	SendComponentSelect(title, oneID, botID string, templates []OneChatRequestElementSendComponent) int
}
type OneChatRequestElementSendComponent struct {
	Image  string          `json:"image"`
	Title  string          `json:"title"`
	Detail string          `json:"detail"`
	Choice []OneChatChoice `json:"choice"`
}

type OneChatChoice struct {
	Url   string `json:"url"`
	Sign  string `json:"sign"` // ? true | false
	Size  string `json:"size"` // ? tall | full
	Type  string `json:"type"` // ? webview | link
	Label string `json:"label"`
}

type OneChatComponentResponse struct {
	Status string `json:"status"`
}
