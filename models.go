package yandexgpt

var (
	GPTModelLite = GPTModel{ModelName: "yandexgpt-lite/latest"}
	GPTModelPRO  = GPTModel{ModelName: "yandexgpt/latest"}

	API_KEY = KeyType{"Api-Key"}
	Bearer  = KeyType{"Bearer"}

	RoleUser      = RoleModel{RoleName: "user"}
	RoleAssistant = RoleModel{RoleName: "assistant"}
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
