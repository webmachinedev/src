package functions

import (
	"os"

	"github.com/mikerybka/github"
)

func UpdateWebmachinedevFrontend() error {
	functions, err := AllFunctions()
	if err != nil {
		return err
	}

	var files map[string]string

	// TODO regenerate /go.mod and /go.sum

	for _, function := range functions {
		vercelfunctionsrc, err := GenerateVercelFunction(function)
		if err != nil {
			return err
		}
		files["api/"+function.ID+"/index.go"] = vercelfunctionsrc

		javascriptfunctionsrc, err := GenerateJavascriptFunction(function)
		if err != nil {
			return err
		}
		files["functions/"+function.ID+"/index.js"] = javascriptfunctionsrc
	}

	// TODO
	// foreach type
	// generate /pages/:type/index.js
	// generate /pages/:type/[id].js
	// generate /pages/:type/new.js
	// generate /pages/:type/[id]/edit.js ???

	return github.WriteFiles("mikerybka", "webmachine.dev", "main", files, "automated update", os.Getenv("GITHUB_KEY"))
}
