package functions

import (
	"io/fs"

	"github.com/webmachinedev/src/types"
)

func AllFunctions(fsys fs.FS) ([]types.Function, error) {
	maxFunctionFileBytes := 20_000

	matches, err := fs.Glob(fsys, "functions/*")
	if err != nil {
		return nil, err
	}

	for _, filename := range matches {
		file, err := fsys.Open(filename)
		if err != nil {
			return nil, err
		}

		bytes := make([]byte, maxFunctionFileBytes)

		_, err = file.Read(bytes)
		if err != nil {
			return nil, err
		}


	}
	pkg := GetPackage("github.com/webmachinedev/functions")
	return pkg.Functions
}
