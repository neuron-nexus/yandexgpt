package template

import (
	"github.com/neuron-nexus/yandexgpt/v2"
	"github.com/neuron-nexus/yandexgpt/v2/internal/models"
)

type Template struct {
	Role yandexgpt.RoleModel
	Text string
}

func (t *Template) ToMessage() yandexgpt.GPTMessage {
	return yandexgpt.GPTMessage{
		Role: t.Role,
		Text: t.Text,
	}
}

func (t *Template) ToRawMessage() models.Message {
	return models.Message{
		Role: t.Role.String(),
		Text: t.Text,
	}
}

func (t *Template) RoleFromString(role string) yandexgpt.RoleModel {
	switch role {
	case "user":
		return yandexgpt.RoleUser
	case "assistant":
		return yandexgpt.RoleAssistant
	default:
		return yandexgpt.RoleUser
	}
}
