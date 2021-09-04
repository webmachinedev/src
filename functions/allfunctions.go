package functions

import "github.com/webmachinedev/src/types"

func AllFunctions() map[string]types.Function {
	pkg := GetPackage("github.com/webmachinedev/functions")
	return pkg.Functions
}
