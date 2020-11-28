package util

import (
	"sync/atomic"
)

type IDGenerator struct {
	next uint64
}

func (i *IDGenerator) New() uint64 {
	return atomic.AddUint64(&i.next, 1)
}
