package types

type HandlerErrorResponse struct {
	Type       string      `json:"type"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Error      interface{} `json:"error"`
}
