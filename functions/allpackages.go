package functions

import "github.com/webmachinedev/src/types"

func AllPackages() []types.Package {
	pkgids := []string{
		"github.com/webmachinedev/src/functions",
		"github.com/webmachinedev/src/types",
	}

	var pkgs []types.Package
	for _, pkgid := range pkgids {
		pkg := GetPackage(pkgid)
		pkgs = append(pkgs, pkg)
	}

	return pkgs
}
