package id

import (
	"sync/atomic"
)

type Gen struct {
	next uint64
}

func (g *Gen) ID() uint64 {
	return atomic.AddUint64(&g.next, 1)
}
