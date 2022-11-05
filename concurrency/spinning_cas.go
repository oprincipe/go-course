package concurrency

import (
	"runtime"
	"sync/atomic"
)

type Spinlock struct {
	state *int32
}

const free = int32(0)

func (l *Spinlock) Lock() {
	// compare and swap over the state and try to look for free
	// if it's free we'll set to 30, if it's not we just continue and nothing will happen
	// NOTE: for this type of loop always use runtime.Gosched() otherwise your program
	// will lock up a single thread, so if you have on one cpu this will lock up
	for !atomic.CompareAndSwapInt32(l.state, free, 30) { // 30 or any other value but 0
		runtime.Gosched() // Poke the scheduler
	}
}

func (l *Spinlock) Unlock() {
	atomic.StoreInt32(l.state, free) // Once atomic, always atomic!
}
