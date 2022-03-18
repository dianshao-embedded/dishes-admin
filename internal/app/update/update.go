package update

import "github.com/google/wire"

var UpdateSet = wire.NewSet(wire.Struct(new(Update), "*"))

type Update struct {
	SuccessChan map[uint64](chan struct{})
	FailedChan  map[uint64](chan struct{})
}
