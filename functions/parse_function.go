package functions

import (
	"io"

	"github.com/webmachinedev/src/types"
)

func ParseFunction(r io.Reader)( types.Function, error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return types.Function{}, err
	}

}
