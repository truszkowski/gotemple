package temple

import (
	"bytes"
)

func ExampleUseExtend(name string, data interface{}) (string, error) {
	tmplA := `{{ define "a" }}this is "a" {{ . }}{{ end }}`
	tmplB := `{{ define "b" }}this is "b" {{ . }}{{ end }}`
	tmplC := `{{ define "c" }}this is "c" {{ . }}{{ end }}`
	tmplMain := `{{ define "main" }}this is main, {{ extends .Name .Data }}{{ end }}`

	tmpl := NewText("tmpl")
	tmpl.Parse(tmplA)
	tmpl.Parse(tmplB)
	tmpl.Parse(tmplC)
	tmpl.Parse(tmplMain)

	v := struct {
		Name string
		Data interface{}
	}{name, data}

	buf := bytes.NewBuffer([]byte{})
	err := tmpl.ExecuteTemplate(buf, "main", &v)

	return buf.String(), err
}
