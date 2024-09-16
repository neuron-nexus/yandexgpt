package dialog

import (
	"fmt"
	endpoint "github.com/neuron-nexus/yandexgpt/internal/endpoint/app/sync"
	model "github.com/neuron-nexus/yandexgpt/internal/models/sync"
	"github.com/neuron-nexus/yandexgpt/pkg/models/gpt"
	"github.com/neuron-nexus/yandexgpt/pkg/models/role"
	"strconv"
)

const (
	API_KEY string = "Api-Key"
	Bearer  string = "Bearer"
)

type YandexGPTSyncApp struct {
	App           *endpoint.App
	SystemMessage model.Message
	Messages      []model.Message

	LastResponse model.Response
}

func New(Key string, KeyType string, StorageID string) (*YandexGPTSyncApp, error) {

	if KeyType != API_KEY && KeyType != Bearer {
		return nil, fmt.Errorf("invalid auth key type. Supported types: %s, %s", API_KEY, Bearer)
	}

	app := endpoint.New()
	app.InitCredential(Key, KeyType)
	app.InitStorageID(StorageID)
	app.InitMaxTokens("2000")

	Messages := make([]model.Message, 0, 10)

	return &YandexGPTSyncApp{
		App:      app,
		Messages: Messages,
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
	p.App.InitTemperature(strconv.FormatFloat(temperature, 'g', 1, 64))
}

func (p *YandexGPTSyncApp) SetModel(modelName gpt.Model) error {
	if modelName.String() != gpt.PRO.String() && modelName.String() != gpt.Lite.String() {
		return fmt.Errorf("invalid model name. Supported models: %s, %s", gpt.PRO.String(), gpt.Lite.String())
	}

	p.App.InitModel(fmt.Sprintf("gpt://%s/%s", p.App.Credential.StorageID, modelName.String()))
	return nil
}

func (p *YandexGPTSyncApp) SendRequest() (model.Response, error) {

	checkError := p.check()

	if checkError != nil {
		return model.Response{}, checkError
	}

	messages := make([]model.Message, len(p.Messages)+1)
	messages[0] = p.SystemMessage
	messages = append(messages, p.Messages...)

	res, err := p.App.SendRequest(messages...)

	// TODO: check token-length of all messages and delete part of them to make length less than 5000 tokens

	if err == nil {
		p.LastResponse = res
		_ = p.AddMessage(role.Assistant, res.Result.Alternatives[0].Message.Text)
	}

	return res, err
}
