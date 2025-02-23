package cpcp

type ResponseType string

const (
	ResError ResponseType = "e"
	ResOK    ResponseType = "o"
)

type Request struct {
	ID      string `json:"i"`
	Payload string `json:"p"`
}

type Response struct {
	ID      string       `json:"i"`
	Type    ResponseType `json:"t"`
	Payload string       `json:"p"`
}
