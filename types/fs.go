package types

import "io/fs"

type FS interface {
	fs.FS
	WriteFiles(files map[string][]byte) error
}
