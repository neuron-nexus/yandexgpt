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

func (p *YandexGPTSyncApp) SetSystemPrompt(prompt string) {
	p.SystemMessage = model.Message{
		Role: "system",
		Text: prompt,
	}
}

func (p *YandexGPTSyncApp) SetSingleMessage(text string) {
	message := model.Message{
		Role: role.User.String(),
		Text: text,
	}
	p.Message = message
}

func (p *YandexGPTSyncApp) SetTemperature(temperature float64) {
	p.App.InitTemperature(temperature)
}

func (p *YandexGPTSyncApp) SetModel(modelName gpt.Model) {
	p.App.InitModel(fmt.Sprintf("gpt://%s/%s", p.App.Credential.StorageID, modelName.String()))
}

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
