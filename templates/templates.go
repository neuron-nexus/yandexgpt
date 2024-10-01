package templates

import (
	"encoding/csv"
	"os"

	"github.com/neuron-nexus/yandexgpt/v2"
	"github.com/neuron-nexus/yandexgpt/v2/internal/template"
)

type Templates struct {
	Template *(map[string]*template.Template)
}

func NewTemplateList() *Templates {
	return &Templates{
		Template: &(map[string]*template.Template{}),
	}
}

func (t *Templates) Add(name string, template *template.Template) {
	(*t.Template)[name] = template
}

func (t *Templates) Get(name string) *template.Template {
	return (*t.Template)[name]
}

func (t *Templates) GetAll() *(map[string]*template.Template) {
	return t.Template
}

func (t *Templates) ToCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	data := [][]string{
		{"name", "role", "text"},
	}
	for name, template := range *t.Template {
		data = append(data, []string{name, template.Role.String(), template.Text})
	}
	err = writer.WriteAll(data)
	if err != nil {
		return err
	}
	writer.Flush()
	return nil
}

func (t *Templates) FromCSV(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	for _, record := range records {
		t.Add(record[0], &template.Template{
			Role: t.roleFromString(record[1]),
			Text: record[2],
		})
	}
	return nil
}

func (t *Templates) roleFromString(role string) yandexgpt.RoleModel {
	switch role {
	case "user":
		return yandexgpt.RoleUser
	case "assistant":
		return yandexgpt.RoleAssistant
	default:
		return yandexgpt.RoleUser
	}
}
