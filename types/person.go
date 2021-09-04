package types

import "io/fs"

type Person struct {
	Name string
	Files fs.FS
}
