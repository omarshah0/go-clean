package types

type HandlerErrorResponse struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
