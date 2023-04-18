package response

type BaseResponse struct {
	Code 	int    `json:"code"`
	Message string `json:"message"`
	Body 	any    `json:"body"`
}