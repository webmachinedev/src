package functions

import "github.com/webmachinedev/src/types"

func AllTypes() ([]types.Type, error) {
	pkg, err := GetPackage("github.com/webmachinedev/types")
	if err != nil {
		return nil, err
	}
	return pkg.Types, nil
}
