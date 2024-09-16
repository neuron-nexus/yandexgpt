package dialog

import (
	"fmt"
	endpoint "github.com/neuron-nexus/yandexgpt/internal/endpoint/app/sync"
	model "github.com/neuron-nexus/yandexgpt/internal/models/sync"
	"github.com/neuron-nexus/yandexgpt/pkg/models/gpt"
	"github.com/neuron-nexus/yandexgpt/pkg/models/key"
	"github.com/neuron-nexus/yandexgpt/pkg/models/role"
)

type YandexGPTSyncApp struct {
	App           *endpoint.App
	SystemMessage model.Message
	Message       model.Message

	Response model.Response
}

// New creates a new instance of YandexGPTSyncApp with the provided credentials and storage ID.
// It initializes the Yandex GPT endpoint with the given parameters and sets the maximum tokens to 2000.
//
//	Key: The API key for authentication.
//	KeyType: The type of the API key.
//	StorageID: The storage ID for the Yandex GPT model.
//
//	Returns a pointer to a new YandexGPTSyncApp instance.
func New(Key string, KeyType key.Type, StorageID string) *YandexGPTSyncApp {

    app := endpoint.New()
    app.InitCredential(Key, KeyType.String())
    app.InitStorageID(StorageID)
    app.InitMaxTokens(2000)

    return &YandexGPTSyncApp{
        App: app,
    }
}

func (p *YandexGPTSyncApp) ChangeCredentials(Key string, KeyType key.Type) {
	p.App.InitCredential(Key, KeyType.String())
}

// SetSystemPrompt sets the system prompt for the Yandex GPT model.
// The system prompt is a message that provides instructions and context to the GPT model.
// It is used to guide the model's responses and ensure they align with the desired behavior.
//
// Parameters:
//	- prompt: A string representing the system prompt.
//
// Returns:
//	This function does not return any value.
func (p *YandexGPTSyncApp) SetSystemPrompt(prompt string) {
    p.SystemMessage = model.Message{
        Role: "system",
        Text: prompt,
    }
}

// SetSingleMessage sets a single user message for the Yandex GPT model.
// This message will be used in subsequent requests to the model.
//
// Parameters:
//	- text: A string representing the user's message. This message will be sent to the Yandex GPT model
//		along with the system prompt.
//
// Returns:
//	This function does not return any value. The user message is stored in the YandexGPTSyncApp instance
//	and can be used in subsequent SendRequest calls.
func (p *YandexGPTSyncApp) SetSingleMessage(text string) {
    message := model.Message{
        Role: role.User.String(),
        Text: text,
    }
    p.Message = message
}

// SetTemperature sets the temperature for the Yandex GPT model.
// The temperature controls the randomness of the model's output. A higher temperature
// results in more diverse and unpredictable responses, while a lower temperature produces
// more consistent and less creative responses.
//
// Parameters:
//	- temperature: A float64 representing the temperature value. The value should be between 0 and 1.
//		If the provided temperature is less than 0, it will be set to 0. If the provided temperature
//		is greater than 1, it will be set to 1.
//
// Returns:
//	This function does not return any value. The temperature setting is stored in the YandexGPTSyncApp
//	instance and will be used in subsequent SendRequest calls.
func (p *YandexGPTSyncApp) SetTemperature(temperature float64) {
    if temperature < 0 {
        temperature = 0
    }
    if temperature > 1 {
        temperature = 1
    }

    p.App.InitTemperature(temperature)
}

func (p *YandexGPTSyncApp) SetModel(modelName gpt.Model) {
	p.App.InitModel(fmt.Sprintf("gpt://%s/%s", p.App.Credential.StorageID, modelName.String()))
}

// SendRequest sends a request to the Yandex GPT model with the current system prompt and user message.
// It checks for any errors before sending the request and updates the response and error fields accordingly.
//
// Parameters:
//	- None
//
// Returns:
//	- res: A model.Response struct containing the response from the Yandex GPT model.
//	- err: An error if any occurred during the request or response processing. If no error occurred,
//		this will be nil.
func (p *YandexGPTSyncApp) SendRequest() (model.Response, error) {
    checkError := p.check()

    if checkError != nil {
        return model.Response{}, checkError
    }

    res, err := p.App.SendRequest([]model.Message{
        p.SystemMessage,
        {
            Role: role.User.String(),
            Text: p.Message.Text,
        },
    }...)

    if err == nil {
        p.Response = res
    }

    return res, err
}
