package dto

type MessageType string

const (
	BUTTON   MessageType = "button"
	DOCUMENT MessageType = "document"
	TEXT     MessageType = "text"
)

// message request
type WAMessagePayload struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Field string `json:"field"`
	Value struct {
		MessagingProduct string      `json:"messaging_product"`
		Metadata         Metadata    `json:"metadata"`
		Contacts         []Contact   `json:"contacts"`
		Messages         []WAMessage `json:"messages"`
	} `json:"value"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Contact struct {
	Profile struct {
		Name string `json:"name"`
	} `json:"profile"`
	WaID string `json:"wa_id"`
}

type WAMessage struct {
	From      string      `json:"from"`
	ID        string      `json:"id"`
	Timestamp string      `json:"timestamp"`
	Text      Text        `json:"text"`
	Type      MessageType `json:"type"` // only can be - button, document, text
	Document  *Document   `json:"document"`
	Button    *Button     `json:"button"`
}

type Text struct {
	Body string `json:"body"`
}
type Document struct {
	Filename         string `json:"filename"`
	MimeType         string `json:"mime_type"`
	Sha256           string `json:"sha256"`
	ID               string `json:"id"`
	MessagingProduct string `json:"messaging_product"`
}

type Button struct {
	Type    string  `json:"type"`
	Text    string  `json:"text"`
	Payload *string `json:"payload"`
	URL     *string `json:"url"`
}

type WabaRequest struct {
}
