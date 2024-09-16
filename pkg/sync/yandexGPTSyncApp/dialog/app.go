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
	Messages      []model.Message

	LastResponse model.Response

	LastMessageTokens int64
}

// New creates a new instance of YandexGPTSyncApp with the provided credentials and storage ID.
// It initializes the app with the given key, key type, storage ID, and sets the maximum tokens to 2000.
// It also creates a new slice for storing messages with a capacity of 10.
//
// Key: The API key for authentication.
// KeyType: The type of the API key.
// StorageID: The storage ID for the Yandex GPT model.
//
// Returns a pointer to the newly created YandexGPTSyncApp instance.
func New(Key string, KeyType key.Type, StorageID string) *YandexGPTSyncApp {

    app := endpoint.New()
    app.InitCredential(Key, KeyType.String())
    app.InitStorageID(StorageID)
    app.InitMaxTokens(2000)

    Messages := make([]model.Message, 0, 10)

    return &YandexGPTSyncApp{
        App:               app,
        Messages:          Messages,
        LastMessageTokens: 0,
    }
}

func (p *YandexGPTSyncApp) ChangeCredentials(Key string, KeyType key.Type) {
	p.App.InitCredential(Key, KeyType.String())
}

// SetSystemPrompt sets the system prompt for the Yandex GPT model.
// The system prompt is a special instruction that provides the model with context and guidelines for generating responses.
//
// Parameters:
//	- prompt: A string representing the system prompt.
//
// Returns:
// This function does not return any value.
func (p *YandexGPTSyncApp) SetSystemPrompt(prompt string) {
    p.SystemMessage = model.Message{
        Role: "system",
        Text: prompt,
    }
}

// AddMessage adds a new message to the YandexGPTSyncApp's message list.
// The message is validated before being added to ensure it has a valid role.
//
// Parameters:
//	- roleName: The role of the message (e.g., role.User, role.Assistant).
//	- text: The content of the message.
//
// Returns:
//	- An error if the message validation fails, otherwise nil.
func (p *YandexGPTSyncApp) AddMessage(roleName role.Model, text string) error {

    validationError := p.validateMessage(roleName)

    if validationError != nil {
        return validationError
    }

    message := model.Message{
        Role: roleName.String(),
        Text: text,
    }
    p.Messages = append(p.Messages, message)

    return nil
}

// SetTemperature sets the temperature for the Yandex GPT model.
// The temperature controls the randomness of the model's output. A higher temperature
// results in more diverse and unpredictable responses, while a lower temperature produces
// more consistent and less creative responses.
//
// Parameters:
//	- temperature: A float64 representing the temperature value. The value should be
//		between 0 and 1 (inclusive). If the provided temperature is less than 0, it is
//		set to 0. If the provided temperature is greater than 1, it is set to 1.
//
// Returns:
//	This function does not return any value.
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

// GetLastText retrieves the text from the last response generated by the Yandex GPT model.
//
// This function checks if the last response is empty. If it is, it returns an empty string and an error indicating that the last response is empty.
// If the last response is not empty, it retrieves the text from the first alternative in the response and returns it along with a nil error.
//
// Returns:
// - A string representing the text from the last response.
// - An error if the last response is empty.
func (p *YandexGPTSyncApp) GetLastText() (string, error) {
    if len(p.LastResponse.Result.Alternatives) == 0 {
        return "", fmt.Errorf("last response is empty")
    }
    return p.LastResponse.Result.Alternatives[0].Message.Text, nil
}

// SendRequest sends a request to the Yandex GPT model with the current system message and messages.
// It first checks if the YandexGPTSyncApp instance is valid before proceeding.
// Then, it creates a new slice of messages, appending the system message and all existing messages.
// After that, it sends the request to the Yandex GPT model using the App's SendRequest method.
// If the request is successful, it updates the LastResponse field with the received response,
// adds the generated response text as a new message with the Assistant role, and returns the response along with a nil error.
// If any error occurs during the request or response handling, it returns an empty Response and the error.
func (p *YandexGPTSyncApp) SendRequest() (model.Response, error) {

    checkError := p.check()

    if checkError != nil {
        return model.Response{}, checkError
    }

    messages := make([]model.Message, 0, len(p.Messages)+1)
    messages = append(messages, p.SystemMessage)
    messages = append(messages, p.Messages...)

    res, err := p.App.SendRequest(messages...)

    if err != nil {
        return res, err
    }

    p.LastResponse = res
    _ = p.AddMessage(role.Assistant, res.Result.Alternatives[0].Message.Text)

    return res, err
}
