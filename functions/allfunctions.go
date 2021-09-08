package functions

import (
	"github.com/webmachinedev/src/types"
)

var SIZE_OF_LARGEST_FUNCTION_FILE = 20_000

//go:embed *
var f embed.FS

func AllFunctions() (functions []types.Function, err error) {

	dirEntries, err := f.ReadDir(".")

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

	return functions, err
}
