package file

import "github.com/google/wire"

var FileServerSet = wire.NewSet(wire.Struct(new(FileServer)), "*")

type FileServer struct {
}
