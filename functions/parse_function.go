package functions

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"strings"

	"github.com/webmachinedev/src/types"
)

func ParseFunction(r io.Reader)(function types.Function, err error) {
	bytes, err := io.ReadAll(r)
	if err != nil {
		return function, err
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", bytes, 0)
	if err != nil {
		return function, err
	}

	str := string(bytes)

	for _, obj := range f.Scope.Objects {
		if decl, ok :=  obj.Decl.(*ast.FuncDecl) ; ok {
			function.Name = obj.Name
			function.GoPackageName = strings.ToLower(obj.Name)
	
			for _, input := range decl.Type.Params.List {
				t := str[input.Type.Pos()-1:input.Type.End()-1]
				var name string
				if len(input.Names) == 0 {
					name = GetNameFromType(t)
				} else {
					name = input.Names[0].Name
				}
				function.Inputs = append(function.Inputs, struct{
					Name string
					Type string
				}{
					Name: name,
					Type: t,
				})
			}
	
			for _, output := range decl.Type.Results.List {
				t := str[output.Type.Pos()-1:output.Type.End()-1]
				var name string
				if len(output.Names) == 0 {
					name = GetNameFromType(t)
				} else {
					name = output.Names[0].Name
				}
				function.Outputs = append(function.Outputs, struct{
					Name string
					Type string
				}{
					Name: name,
					Type: t,
				})
			}
		}
	}
	return function, err
}
