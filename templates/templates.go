package templates

import (
	"github.com/neuron-nexus/yandexgpt/v2"
	"github.com/neuron-nexus/yandexgpt/v2/internal/models"
)

type Template struct {
	Name string
	Role yandexgpt.RoleModel
	Text string
}

func New(Name string, Role yandexgpt.RoleModel, Message string) *Template {
	return &Template{
		Name: Name,
		Role: Role,
		Text: Message,
	}
}

func (t *Template) String() string {
	return t.Text
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
