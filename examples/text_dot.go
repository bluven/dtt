package main

import (
	"dtt/template"
	"os"
)

func main() {
	const text = `
Hello, world

{{- .Name }}
`

	tpl := template.Must(template.New("debug").Parse(text))

	_ = tpl.Execute(os.Stdout, nil)
}
