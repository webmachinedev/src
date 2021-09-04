package functions

import "github.com/webmachinedev/src/types"

func AllTypes() []types.Type {
	pkg := GetPackage("github.com/webmachinedev/types")
	return pkg.Types
}
