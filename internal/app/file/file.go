package file

import "github.com/google/wire"

var FileSet = wire.NewSet(
	FileServerSet,
) // end
