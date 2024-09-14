package yandexGPTSyncApp

import (
	"fmt"
	model "github.com/neuron-nexus/yandexgpt/internal/models/sync"
	"strconv"
)

const (
	API_KEY string = "Api-Key"
	Bearer  string = "Bearer"
	URL     string = "https://llm.api.cloud.yandex.net/foundationModels/v1/completion"
)

type YandexGPTSyncApp struct {
	Key          string
	KeyType      string
	StorageID    string
	SystemPrompt string
}

// New creates a new instance of App with the provided key and key type.
//
// The function accepts two parameters:
//   - Key: A string representing the authentication key.
//   - KeyType: A string representing the type of the authentication key. It can be either "API_KEY" or "Bearer".
//
// The function returns two values:
//   - A pointer to an App instance if the key type is valid.
//   - An error if the key type is invalid. The error message will include the supported key types.
//
// Note: The KeyType should be one of the constants defined in the package: API_KEY or Bearer.
func New(Key string, KeyType string, StorageID string) (*YandexGPTSyncApp, error) {

	if KeyType != API_KEY && KeyType != Bearer {
		return nil, fmt.Errorf("invalid auth key type. Supported types: %s, %s", API_KEY, Bearer)
	}

	return &YandexGPTSyncApp{
		Key:       Key,
		KeyType:   KeyType,
		StorageID: StorageID,
	}, nil
}

// ChangeCredentials updates the authentication key and key type for the YandexGPTSyncApp instance.
//
// The function accepts two parameters:
//   - Key: A string representing the new authentication key.
//   - KeyType: A string representing the new type of the authentication key. It can be either "API_KEY" or "Bearer".
//
// The function returns an error if the provided KeyType is invalid. The error message will include the supported key types.
//
// If the provided KeyType is valid, the function updates the Key and KeyType fields of the YandexGPTSyncApp instance and returns nil.
func (p *YandexGPTSyncApp) ChangeCredentials(Key string, KeyType string) error {
	if KeyType != API_KEY && KeyType != Bearer {
		return fmt.Errorf("invalid auth key type. Supported types: %s, %s", API_KEY, Bearer)
	}

	p.Key = Key
	p.KeyType = KeyType
	return nil
}

// SetSystemPrompt updates the system prompt for the YandexGPTSyncApp instance.
//
// The function accepts a single parameter:
//   - prompt: A string representing the new system prompt. It cannot be an empty string.
//
// The function returns an error if the provided prompt is empty. The error message will indicate that a prompt is required.
//
// If the provided prompt is not empty, the function updates the SystemPrompt field of the YandexGPTSyncApp instance and returns nil.
func (p *YandexGPTSyncApp) SetSystemPrompt(prompt string) error {
	if prompt == "" {
		return fmt.Errorf("invalid prompt. Prompt is required")
	}
	p.SystemPrompt = prompt
	return nil
}

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

	modelUri := fmt.Sprintf("gpt://%s/%s", p.StorageID, modelVersion)
	stringTemperature := strconv.FormatFloat(temperature, 'g', 1, 64)
}
