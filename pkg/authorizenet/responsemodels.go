package authorizenet

type ResponseMessage struct {
	Code string `json:"code"`
	Text string `json:"text"`
}

type ResponseMessages struct {
	ResultCode string            `json:"resultCode"`
	Message    []ResponseMessage `json:"message"`
}

type TransactionResponse struct {
	ResponseCode           string                       `json:"responseCode"`
	TransactionId          string                       `json:"transId"`
	ReferenceTransactionId string                       `json:"refTransID,omitempty"`
	Messages               []TransactionResponseMessage `json:"messages,omitempty"`
}

type TransactionResponseMessage struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type Response struct {
	TransactionResponse TransactionResponse `json:"transactionResponse"`
	RefId               string              `json:"refId"`
	Messages            ResponseMessages    `json:"messages,omitempty"`
}
