package yandexgpt

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	endpoint "github.com/neuron-nexus/yandexgpt/internal/endpoint/app/sync"
	model "github.com/neuron-nexus/yandexgpt/internal/models"
)

type YandexGPTSyncApp struct {
	App           *endpoint.App
	SystemMessage model.Message
	Message       []model.Message

	Response Response

	DebugMode bool
}

func NewYandexGPTSyncApp(Key string, KeyType KeyType, StorageID string, Model GPTModel) *YandexGPTSyncApp {

	app := endpoint.New()
	app.InitCredential(Key, KeyType.String())
	app.InitStorageID(StorageID)
	app.InitMaxTokens(2000)
	app.InitTemperature(0.3)
	app.InitModel(fmt.Sprintf("gpt://%s/%s", StorageID, Model.String()))

	return &YandexGPTSyncApp{
		App:       app,
		DebugMode: false,
	}
}

func (p *YandexGPTSyncApp) ChangeCredentials(Key string, KeyType KeyType) {
	if p.DebugMode {
		log.Println("Initialization new credentials:", KeyType.String(), "/", Key)
	}
	p.App.InitCredential(Key, KeyType.String())
}

func (p *YandexGPTSyncApp) Configure(Parameters ...GPTParameter) error {
	for _, param := range Parameters {
		switch strings.ToLower(param.Name.String()) {
		case "prompt":

			if p.DebugMode {
				log.Println("Initialization prompt:", param.Value)
			}

			if param.Value == "" {
				return fmt.Errorf("empty prompt")
			}
			p.SystemMessage = model.Message{
				Role: "system",
				Text: param.Value,
			}
			continue

		case "temperature":

			if p.DebugMode {
				log.Println("Initialization temperature:", param.Value)
			}

			temperature, err := strconv.ParseFloat(param.Value, 64)
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
			continue

		case "max_tokens":

			if p.DebugMode {
				log.Println("Initialization max_tokens:", param.Value)
			}

			maxTokens, err := strconv.ParseInt(param.Value, 10, 64)
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
			continue

		default:
			return fmt.Errorf("unknown parameter: %s\nuse:\n-%s\n-%s\n-%s", param.Name.String(), "prompt", "temperature", "max_tokens")
		}
	}
	return nil
}

func (p *YandexGPTSyncApp) AddMessage(Message GPTMessage) error {
	if Message.Text == "" {
		return fmt.Errorf("empty message")
	}
	if Message.Role.String() == "" {
		Message.Role = RoleUser
	}
	if Message.Role.String() != "user" && Message.Role.String() != "assistant" {
		return fmt.Errorf("unknown role: %s. Use: \"user\" or \"assistant\"", Message.Role.String())
	}

	if p.DebugMode {
		log.Println("Add message:", Message.Text)
	}

	p.Message = append(p.Message, model.Message{
		Role: Message.Role.String(),
		Text: Message.Text,
	})

	return nil
}

// Unsafe: AddRawMessage is unsafe function. Use AddMessage(Message GPTMessage)
func (p *YandexGPTSyncApp) AddRawMessage(Message model.Message) error {
	if Message.Text == "" {
		return fmt.Errorf("empty message")
	}
	if Message.Role == "" {
		Message.Role = "user"
	}
	if Message.Role != "user" && Message.Role != "assistant" {
		return fmt.Errorf("unknown role: %s. Use: \"user\" or \"assistant\"", Message.Role)
	}

	if p.DebugMode {
		log.Println("Add message:", Message.Text)
	}

	p.Message = append(p.Message, Message)

	return nil
}

func (p *YandexGPTSyncApp) SetMessages(Messages ...GPTMessage) error {
	p.Message = nil

	var messages []model.Message = make([]model.Message, 0, len(Messages))

	if p.DebugMode {
		log.Println("Set messages:", Messages)
	}

	for _, message := range Messages {
		if message.Text == "" {
			return fmt.Errorf("empty message")
		}
		messages = append(messages, model.Message{
			Role: message.Role.String(),
			Text: message.Text,
		})
	}

	p.Message = append(p.Message, messages...)
	return nil
}

func (p *YandexGPTSyncApp) ClearMessages() {
	if p.DebugMode {
		log.Println("Clear messages")
	}
	p.Message = nil
}

func (p *YandexGPTSyncApp) SendRequest() (Response, error) {

	if p.SystemMessage.Text == "" {
		return Response{}, fmt.Errorf("empty prompt. use Configure(\"prompt\", \"YOUR_PROMPT\")")
	}

	messages := []model.Message{}

	messages = append(messages, p.SystemMessage)
	messages = append(messages, p.Message...)

	if p.DebugMode {
		log.Println("Send request:", messages)
	}

	res, err := p.App.SendRequest(messages...)

	if err != nil {
		return Response{}, err
	}

	if len(res.Result.Alternatives) == 0 {
		if p.DebugMode {
			log.Println(res)
		}
		return Response{}, fmt.Errorf("empty response")
	}

	response := Response{
		Result: res.Result,
		Text:   res.Result.Alternatives[0].Message.Text,
	}

	p.Response = response

	return response, err
}
