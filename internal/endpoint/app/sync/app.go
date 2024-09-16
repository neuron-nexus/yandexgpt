package sync

import model "github.com/neuron-nexus/yandexgpt/internal/models/sync"

type App struct {
	Credential struct {
		Key       string
		KeyType   string
		StorageID string
	}

	CompletionOptions struct {
		MaxTokens   string
		Temperature string
	}

	ModelUri string
	Messages []model.Message
}

func New() *App {
	return &App{}
}

func (a *App) InitCredential(Key string, KeyType string) {
	a.Credential.Key = Key
	a.Credential.KeyType = KeyType
}

func (a *App) InitStorageID(StorageID string) {
	a.Credential.StorageID = StorageID
}

func (a *App) InitModel(ModelUri string) {
	a.ModelUri = ModelUri
}

func (a *App) InitMaxTokens(MaxTokens string) {
	a.CompletionOptions.MaxTokens = MaxTokens
}

func (a *App) InitTemperature(Temperature string) {
	a.CompletionOptions.Temperature = Temperature
}

func (a *App) SendRequest(messages ...model.Message) (model.Response, error) {
	Request := model.Request{}

	Request.ModelUri = a.ModelUri

	Request.CompletionOptions.Stream = "false"
	Request.CompletionOptions.MaxTokens = a.CompletionOptions.MaxTokens
	Request.CompletionOptions.Temperature = a.CompletionOptions.Temperature

	Request.Messages = messages

	return a.sendRequestToYandexGPT(&Request)
}