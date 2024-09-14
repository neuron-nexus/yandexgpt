package sync

type Message struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type Request struct {
	ModelUri          string `json:"modelUri"`
	CompletionOptions struct {
		Stream      string `json:"stream"`
		Temperature string `json:"temperature"`
		MaxTokens   string `json:"maxTokens"`
	} `json:"completionOptions"`
	Messages []Message `json:"messages"`
}
