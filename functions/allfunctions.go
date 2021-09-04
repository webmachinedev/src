package functions

import "github.com/webmachinedev/src/types"

func AllFunctions() []types.Function {
	pkg := GetPackage("github.com/webmachinedev/functions")
	return pkg.Functions
}
