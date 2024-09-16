package gpt

var (
	Lite = Model{ModelName: "yandexgpt-lite/latest"}
	PRO  = Model{ModelName: "yandexgpt/latest"}
)

type Model struct {
	ModelName string
}

func (m *Model) String() string {
	return m.ModelName
}
