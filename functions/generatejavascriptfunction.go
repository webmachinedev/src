package functions

import (
	"bytes"
	"text/template"

	"github.com/webmachinedev/src/types"
)

func GenerateJavascriptFunction(function types.Function) (string, error) {
	javascriptfunctiontmpl := `export default async function {{ .ID }}() {
	}
	`
	filetmpl, err := template.New("javascriptfunction").Parse(javascriptfunctiontmpl)
	if err != nil {
		return "", err
	}

	contents := bytes.NewBuffer(nil)
	err = filetmpl.Execute(contents, function)
	if err != nil {
		return "", err
	}

	return contents.String(), nil
}