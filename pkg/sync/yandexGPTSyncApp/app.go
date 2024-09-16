package yandexGPTSyncApp

import (
	"fmt"
	endpoint "github.com/neuron-nexus/yandexgpt/internal/endpoint/app/sync"
	model "github.com/neuron-nexus/yandexgpt/internal/models/sync"
	"strconv"
)

const (
	API_KEY string = "Api-Key"
	Bearer  string = "Bearer"
)

type YandexGPTSyncApp struct {
	App           *endpoint.App
	SystemMessage model.Message
}

func New(Key string, KeyType string, StorageID string) (*YandexGPTSyncApp, error) {

	if KeyType != API_KEY && KeyType != Bearer {
		return nil, fmt.Errorf("invalid auth key type. Supported types: %s, %s", API_KEY, Bearer)
	}

	app := endpoint.New()
	app.InitCredential(Key, KeyType)
	app.InitStorageID(StorageID)
	app.InitMaxTokens("2000")

	return &YandexGPTSyncApp{
		App: app,
	}, nil
}

func (p *YandexGPTSyncApp) ChangeCredentials(Key string, KeyType string) error {
	if KeyType != API_KEY && KeyType != Bearer {
		return fmt.Errorf("invalid auth key type. Supported types: %s, %s", API_KEY, Bearer)
	}

	p.App.InitCredential(Key, KeyType)

	return nil
}

func (p *YandexGPTSyncApp) SetSystemPrompt(prompt string) error {
	if prompt == "" {
		return fmt.Errorf("invalid prompt. Prompt is required")
	}

	p.SystemMessage = model.Message{
		Role: "system",
		Text: prompt,
	}

	return nil
}

// SendRequest sends a request to the Yandex GPT model with the provided parameters.
//
// Parameters:
//   - modelVersion: A string representing the version of the GPT model to be used.
//     It can be either "pro" or "lite". If "pro" is provided, it will be replaced with "yandexgpt/latest".
//     If "lite" is provided, it will be replaced with "yandexgpt-lite/latest".
//   - temperature: A float64 representing the randomness of the model's output.
//     A higher value will make the output more diverse, while a lower value will make it more deterministic.
//   - messages: A variadic parameter of type model.Message representing the messages to be sent to the model.
//
// Return:
//   - A model.Response representing the response from the model.
//   - An error if any validation error occurs or if there is an error sending the request.
func (p *YandexGPTSyncApp) SendRequest(modelVersion string, temperature float64, messages ...model.Message) (model.Response, error) {

	validationError := p.validateRequestInput(modelVersion, temperature, messages...)
	if validationError != nil {
		return model.Response{}, validationError
	}

	if modelVersion == "pro" {
		modelVersion = "yandexgpt/latest"
	} else {
		modelVersion = "yandexgpt-lite/latest"
	}

	p.App.InitModel(fmt.Sprintf("gpt://%s/%s", p.App.Credential.StorageID, modelVersion))
	p.App.InitTemperature(strconv.FormatFloat(temperature, 'g', 1, 64))
	return p.App.SendRequest(messages...)
}
