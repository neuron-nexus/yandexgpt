package validation

import (
	"fmt"
	model "github.com/neuron-nexus/yandexgpt/internal/models/sync"
)

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateMessages(messages ...model.Message) error {
	for _, msg := range messages {
		if msg.Role == "" || msg.Text == "" {
			return fmt.Errorf("invalid message. Role and text are required")
		}
		if msg.Role != "assistant" && msg.Role != "user" {
			return fmt.Errorf("invalid message. Role should be either 'assistant' or 'user'")
		}
	}
	return nil
}
