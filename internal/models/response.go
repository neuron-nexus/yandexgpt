package models

type Alternative struct {
	Message Message `json:"message"`
	Status  string  `json:"status"`
}

type Result struct {
	Alternatives []Alternative `json:"alternatives"`
	Usage        struct {
		InputTextTokens  string `json:"input_text_tokens"`
		CompletionTokens string `json:"completionTokens"`
		TotalTokens      string `json:"totalTokens"`
	} `json:"usage"`
	ModelVersion string `json:"modelVersion"`
}

type Response struct {
	Result Result `json:"result"`
}
