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

func (p *YandexGPTSyncApp) SetSystemPrompt(prompt string) {
	p.SystemMessage = model.Message{
		Role: "system",
		Text: prompt,
	}
}

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

func (p *YandexGPTSyncApp) GetLastText() (string, error) {
	if len(p.LastResponse.Result.Alternatives) == 0 {
		return "", fmt.Errorf("last response is empty")
	}
	return p.LastResponse.Result.Alternatives[0].Message.Text, nil
}

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
