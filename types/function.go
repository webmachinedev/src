package types

type Function struct {
	ID string
	Name string
	GoPackageName string
	Inputs []Field
	Outputs []Field
	Body string
}

// func (f *Function) GoPackageName() string {
// 	validchars := "abcdefghijklmnopqrstuvwxyz"
// 	name := ""
// 	for _, char := range strings.ToLower(f.Name) {
// 		if strings.ContainsRune(validchars, char) {
// 			name += string(char)
// 		}
// 	}
// 	return name
// }
