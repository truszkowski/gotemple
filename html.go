package tmpl

import (
	"bytes"
	"html/template"
)

// Creates HTML template with predefined functions
func NewHtml(name string) *template.Template {
	tmpl := template.New(name)

	extends := template.FuncMap(map[string]interface{}{
		"extends": func(name string, data interface{}) (html template.HTML, err error) {
			buf := bytes.NewBuffer([]byte{})

			if err = tmpl.ExecuteTemplate(buf, name, data); err == nil {
				html = template.HTML(buf.String())
			}

			return
		},
	})

	return tmpl.Funcs(template.FuncMap(coreFuncs)).Funcs(extends)
}
