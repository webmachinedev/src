package functions

import "github.com/webmachinedev/src/types"

func AllFunctions(caller types.Person) []types.Function {
	pkg := GetPackage("github.com/webmachinedev/functions")
	return pkg.Functions
}
