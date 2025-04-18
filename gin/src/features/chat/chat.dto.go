package chat

type ChatRequestData struct{
	ToUser string `json:"toUser"`
	Message string `json:"message"`
}


type ChatResponseData struct{
	FromUser string `json:"fromUser"`
	Message string `json:"message"`
}


type ChatErrorResponse struct{
	Message string `json:"message"`
}
