package concurrency

import (
	"runtime"
	"sync/atomic"
)

type TicketStore struct {
	ticket *uint64
	done   *uint64
	slots  []string // for simplicity imagine this to be infinite (usually could be some stream buffer)
}

func (ts *TicketStore) Put(s string) {
	t := atomic.AddUint64(ts.ticket, 1) - 1 // draw a ticket (-1 because we get the value after the operation)
	ts.slots[t] = s                         // store your data

	// this block is keeping us from being weight free in this function
	for !atomic.CompareAndSwapUint64(ts.done, t, t+1) { // increase done
		runtime.Gosched()
	}
}

func (ts *TicketStore) GetDone() []string {
	return ts.slots[:atomic.LoadUint64(ts.done)+1] // read up to done
}
