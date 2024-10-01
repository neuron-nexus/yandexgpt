package yandexgpt

import (
	"fmt"
	"strconv"
	"strings"

	endpoint "github.com/neuron-nexus/yandexgpt/internal/endpoint/app/sync"
	model "github.com/neuron-nexus/yandexgpt/internal/models"
)

type YandexGPTSyncApp struct {
	App           *endpoint.App
	SystemMessage model.Message
	Message       []model.Message

	Response model.Response
}

func NewYandexGPTSyncApp(Key string, KeyType KeyType, StorageID string, Model GPTModel) *YandexGPTSyncApp {

	app := endpoint.New()
	app.InitCredential(Key, KeyType.String())
	app.InitStorageID(StorageID)
	app.InitMaxTokens(2000)
	app.InitTemperature(0.3)
	app.InitModel(fmt.Sprintf("gpt://%s/%s", app.Credential.StorageID, Model.String()))

	return &YandexGPTSyncApp{
		App: app,
	}
}

func (p *YandexGPTSyncApp) ChangeCredentials(Key string, KeyType KeyType) {
	p.App.InitCredential(Key, KeyType.String())
}

func (p *YandexGPTSyncApp) Configure(Parameter, Value string) error {
	switch strings.ToLower(Parameter) {
	case "prompt":
		if Value == "" {
			return fmt.Errorf("empty prompt")
		}
		p.SystemMessage = model.Message{
			Role: "system",
			Text: Value,
		}
		return nil

	case "temperature":
		temperature, err := strconv.ParseFloat(Value, 64)
		if err != nil {
			return err
		}
		if temperature < 0 {
			temperature = 0
		}
		if temperature > 1 {
			temperature = 1
		}
		p.App.InitTemperature(temperature)
		return nil

	case "max_tokens":
		maxTokens, err := strconv.ParseInt(Value, 10, 64)
		if err != nil {
			return err
		}
		if maxTokens < 0 {
			maxTokens = 0
		}
		if maxTokens > 2000 {
			maxTokens = 2000
		}
		p.App.InitMaxTokens(maxTokens)
		return nil

	default:
		return fmt.Errorf("unknown parameter: %s\nuse:\n-%s\n-%s\n-%s", Parameter, "prompt", "temperature", "max_tokens")
	}
}

func (p *YandexGPTSyncApp) AddMessage(Role RoleModel, Text string) {
	p.Message = append(p.Message, model.Message{
		Role: Role.String(),
		Text: Text,
	})
}

func (p *YandexGPTSyncApp) SetMessages(Messages ...model.Message) {
	p.Message = nil
	p.Message = append(p.Message, Messages...)
}

func (p *YandexGPTSyncApp) ClearMessages() {
	p.Message = nil
}

func (p *YandexGPTSyncApp) SendRequest() (model.Response, error) {

	if p.SystemMessage.Text == "" {
		return model.Response{}, fmt.Errorf("empty prompt. use Configure(\"prompt\", \"YOUR_PROMPT\")")
	}

	messages := []model.Message{}

	messages = append(messages, p.SystemMessage)
	messages = append(messages, p.Message...)

	res, err := p.App.SendRequest(messages...)

	if err == nil {
		p.Response = res
	}

	return res, err
}
