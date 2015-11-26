package temple

import (
	"bytes"
	"text/template"
)

// Creates TEXT template with predefined functions
func NewText(name string) *template.Template {
	tmpl := template.New(name)

	extends := template.FuncMap(map[string]interface{}{
		"extends": func(name string, data interface{}) (text string, err error) {
			buf := bytes.NewBuffer([]byte{})

			if err = tmpl.ExecuteTemplate(buf, name, data); err == nil {
				text = buf.String()
			}

			return
		},
	})

	return tmpl.Funcs(template.FuncMap(coreFuncs)).Funcs(extends)
}
