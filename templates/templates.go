package templates

type Template struct {
	TemplateString string
}

func New(templateString string) *Template {
	return &Template{
		TemplateString: templateString,
	}
}

func (t *Template) String() string {
	return t.TemplateString
}
