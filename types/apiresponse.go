package types

type ErrorResponse struct {
	Type    string      `json:"type"`
	Trace   string      `json:"trace"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
