package templates

import (
	"html/template"
	"io"
	"strings"
)

type LayoutTemplateProcessor struct {
}

func (l *LayoutTemplateProcessor) ExecTemplate(w io.Writer, name string, data any) error {
	var sb strings.Builder
	layoutName := ""
	localTemplate := getTemplates()
	localTemplate.Funcs(map[string]any{
		"body":   insertBodyWrapper(&sb),
		"layout": insertLayoutWrapper(&layoutName),
	})
	err := localTemplate.ExecuteTemplate(&sb, name, data)
	if err != nil {
		return err
	}
	if layoutName != "" {
		return localTemplate.ExecuteTemplate(w, layoutName, data)
	}
	_, err = io.WriteString(w, sb.String())
	if err != nil {
		return err
	}

	return nil
}

func insertBodyWrapper(s *strings.Builder) func() template.HTML {
	return func() template.HTML {
		return template.HTML(s.String())
	}
}

func insertLayoutWrapper(s *string) func(string) string {
	return func(layout string) string {
		*s = layout
		return ""
	}
}

var getTemplates func() *template.Template
