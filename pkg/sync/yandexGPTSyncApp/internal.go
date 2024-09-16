package yandexGPTSyncApp

import (
	"fmt"
	model "github.com/neuron-nexus/yandexgpt/internal/models/sync"
)

func (p *YandexGPTSyncApp) validateRequestInput(modelVersion string, temperature float64, messages ...model.Message) error {
	if modelVersion != "pro" && modelVersion != "lite" {
		return fmt.Errorf("invalid model version. Supported versions: pro, lite")
	}

	if temperature < 0 || temperature > 1 {
		return fmt.Errorf("invalid temperature. Temperature should be between 0 and 1")
	}

	if p.SystemMessage.Text == "" {
		return fmt.Errorf("invalid system prompt. System prompt is required. Use SetSystemPrompt(prompt syting)")
	}

	return nil
}
