package functions

import (
	"github.com/webmachinedev/src/types"
)

var SIZE_OF_LARGEST_FUNCTION_FILE = 20_000

func AllFunctions() ([]types.Function, error) {
	pkg, err := GetPackage("github.com/webmachinedev/types")
	if err != nil {
		return nil, err
	}
	return pkg.Functions, nil
	// //go:embed *
	// var f embed.FS

	// dirEntries, err := f.ReadDir(".")
	// if err != nil {
	// 	return nil, err
	// }

	// var functions []types.Function

	// for _, dirEntry := range dirEntries {
	// 	file, err := f.Open(dirEntry.Name())
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	function, err := ParseFunction(file)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	functions = append(functions, function)
	// }

	// return functions, nil
}
