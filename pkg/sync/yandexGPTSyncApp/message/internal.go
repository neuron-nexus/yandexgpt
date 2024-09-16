package dialog

import (
	"fmt"
)

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
			err = fmt.Errorf("invalid model URI. Set model URI before using the app. Use SetModel function\n")
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

	return err
}
