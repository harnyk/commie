package cpcp

import "encoding/json"

type ErrorResponse struct {
	ErrorPayload string `json:"error"`
}

func NewErrorResponse(payload string) *ErrorResponse {
	return &ErrorResponse{ErrorPayload: payload}
}

func (e *ErrorResponse) Error() string {
	return e.ErrorPayload
}
func (e *ErrorResponse) Payload(payload any) error {
	return json.Unmarshal([]byte(e.ErrorPayload), payload)
}
