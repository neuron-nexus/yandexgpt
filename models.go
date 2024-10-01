package yandexgpt

import model "github.com/neuron-nexus/yandexgpt/v2/internal/models"

var (
	GPTModelLite = GPTModel{ModelName: "yandexgpt-lite/latest"}
	GPTModelPRO  = GPTModel{ModelName: "yandexgpt/latest"}

	API_KEY = KeyType{"Api-Key"}
	Bearer  = KeyType{"Bearer"}

	RoleUser      = RoleModel{RoleName: "user"}
	RoleAssistant = RoleModel{RoleName: "assistant"}

	ParameterTemperature = GPTParameterName{Name: "temperature"}
	ParameterPrompt      = GPTParameterName{Name: "prompt"}
	ParameterMaxTokens   = GPTParameterName{Name: "max_tokens"}
)

// GPTModel - models of yandexGPT
type GPTModel struct {
	ModelName string
}

func (t *GPTModel) String() string {
	return t.ModelName
}

// KeyType - model of key types
type KeyType struct {
	KeyType string
}

func (t *KeyType) String() string {
	return t.KeyType
}

// RoleModel - model of roles
type RoleModel struct {
	RoleName string
}

func (m *RoleModel) String() string {
	return m.RoleName
}

type GPTMessage struct {
	Role RoleModel
	Text string
}

type GPTParameterName struct {
	Name string
}

func (n *GPTParameterName) String() string {
	return n.Name
}

type GPTParameter struct {
	Name  GPTParameterName
	Value string
}

type Response struct {
	Result model.Result
	Text   string
}
