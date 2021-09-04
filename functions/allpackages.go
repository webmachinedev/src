package functions

import (
	"io/fs"

	"github.com/webmachinedev/src/types"
)

func AllPackages(fs fs.FS) []types.Package {
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
