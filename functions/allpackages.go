package functions

import "github.com/webmachinedev/src/types"

func AllPackages() map[string]*types.Package {
	pkgids := []string{
		"github.com/webmachinedev/functions",
		"github.com/webmachinedev/src/types",
	}

	for _, pkgid := range pkgids {
		pkg := GetPackage(pkgid)
		
	}
}
