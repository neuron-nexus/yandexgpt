package models

type Alternative struct {
	Message Message `json:"message"`
	Status  string  `json:"status"`
}

type Result struct {
	Alternatives []Alternative `json:"alternatives"`
	Usage        struct {
		InputTextTokens  int64 `json:"input_text_tokens"`
		CompletionTokens int64 `json:"completionTokens"`
		TotalTokens      int64 `json:"totalTokens"`
	} `json:"usage"`
	ModelVersion string `json:"modelVersion"`
}

type Response struct {
	Result Result `json:"result"`
}
