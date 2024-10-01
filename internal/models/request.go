package models

type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type Request struct {
	ModelUri          string `json:"modelUri"`
	CompletionOptions struct {
		Stream      bool    `json:"stream"`
		Temperature float64 `json:"temperature"`
		MaxTokens   int64   `json:"maxTokens"`
	} `json:"completionOptions"`
	Messages []Message `json:"messages"`
}
