package functions

func GetNameFromType(t string) (name string) {
	switch t {
	case "error":
		return "err"
	case "string":
		return "str"
	case "[]string":
		return "ids"
	case "[]types.Type":
		return "types"
	case "types.Package":
		return "package"
	default:
		panic("no name for type " + t)
	}
}
