package dialog

import (
	"fmt"
	"github.com/neuron-nexus/yandexgpt/pkg/models/role"
)

func (p *YandexGPTSyncApp) validateMessage(roleName role.Model) error {
	if p.SystemMessage.Text == "" {
		return fmt.Errorf("invalid prompt. Set system prompt before adding messages. Use SetSystemPrompt function")
	}

	if roleName.String() != "user" && roleName.String() != "assistant" {
		return fmt.Errorf("invalid role. Supported roles: %s, %s", role.User.String(), role.Assistant.String())
	}

	if len(p.Messages) == 0 && roleName.String() == "assistant" {
		return fmt.Errorf("invalid role. First message must have \"user\" role")
	}

	if len(p.Messages) > 0 && p.Messages[len(p.Messages)-1].Role == roleName.String() {
		return fmt.Errorf("invalid role. Previous message have the same role")
	}

	return nil
}

func (p *YandexGPTSyncApp) check() error {
	var err error = nil

	if p.App.Credential.Key == "" {
		err = fmt.Errorf("invalid auth key. Set credentials before using the app. Use ChangeCredentials function\n")
	}

	if p.App.Credential.KeyType == "" {
		if err == nil {
			err = fmt.Errorf("invalid auth key type. Set credentials before using the app. Use ChangeCredentials function\n")
		} else {
			err = fmt.Errorf(err.Error() + "invalid auth key type. Set credentials before using the app. Use ChangeCredentials function\n")
		}
	}

	if p.App.Credential.StorageID == "" {
		if err == nil {
			err = fmt.Errorf("invalid storage ID. Set credentials before using the app. Use ChangeCredentials function\n")
		} else {
			err = fmt.Errorf(err.Error() + "invalid storage ID. Set credentials before using the app. Use ChangeCredentials function\n")
		}
	}

	if p.App.CompletionOptions.MaxTokens == 0 {
		if err == nil {
			err = fmt.Errorf("invalid max tokens. Set max tokens before using the app. Use InitMaxTokens function\n")
		} else {
			err = fmt.Errorf(err.Error() + "invalid max tokens. Set max tokens before using the app. Use InitMaxTokens function\n")
		}
	}

	if p.App.CompletionOptions.Temperature == 0 {
		if err == nil {
			err = fmt.Errorf("invalid temperature. Set temperature before using the app. Use InitTemperature function\n")
		} else {
			err = fmt.Errorf(err.Error() + "invalid temperature. Set temperature before using the app. Use InitTemperature function\n")
		}
	}

	if p.App.ModelUri == "" {
		if err == nil {
			err = fmt.Errorf("invalid model URI. Set model URI before using the app. Use InitModel function\n")
		} else {
			err = fmt.Errorf(err.Error() + "invalid model URI. Set model URI before using the app. Use InitModel function\n")
		}
	}

	if p.SystemMessage.Text == "" {
		if err == nil {
			err = fmt.Errorf("invalid prompt. Set system prompt before using the app. Use SetSystemPrompt function\n")
		} else {
			err = fmt.Errorf(err.Error() + "invalid prompt. Set system prompt before using the app. Use SetSystemPrompt function\n")
		}
	}

	if len(p.Messages) == 0 {
		if err == nil {
			err = fmt.Errorf("invalid messages. Add at least one message. Use AddMessage function\n")
		} else {
			err = fmt.Errorf(err.Error() + "invalid messages. Add at least one message. Use AddMessage function\n")
		}
	}

	return err
}
