package concurrency

import "time"

// TryReceive make channels blocking /**
func TryReceive(c <-chan int) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	default: // Processed when c is blocking
		return 0, true, false
	}
}

// TryReceiveWithTimeout make channels non-blocking /**
func TryReceiveWithTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
	select {
	case data, more = <-c:
		return data, more, true
	case <-time.After(duration): // time.After(d) returns a channel
		return 0, true, false
	default: // Processed when c is blocking
		return 0, true, false
	}
}

// Fanout - one stream of data is coming in and multiple come out
func Fanout(In <-chan int, OutA, OutB chan int) {
	// if the channel is closed the loop will discontinue
	for data := range In { // Receive until closed
		select {
		case OutA <- data:
		case OutB <- data:
		}
	}
}

// TurnoutExample1 - one stream of data is coming in and multiple come out
// bad example where you don't know when to quit from the waiting select
func TurnoutExample1(InA, InB <-chan int, OutA, OutB chan int) {
	var data int
	var more bool

	// variable declaration left out for readability
	for {
		select {
		// Receive from first non-blocking
		case data, more = <-InA:
		case data, more = <-InB:
		}

		if !more {
			// ... ?
			return
		}

		select { // Send to first non-blocking
		case OutA <- data:
		case OutB <- data:
		}
	}
}

// TurnoutExample2 - one stream of data is coming in and multiple come out
// good example with quit channel
func TurnoutExample2(Quit <-chan int, InA, InB, OutA, OutB chan int) {
	var data int
	for {
		select {
		case data = <-InA:
		case data = <-InB:

		case <-Quit: // remember: close generates a message
			close(InA) // Actually this is an anti-pattern ... (the receiver should never close)
			close(InB) // ... you can argue that quit acts as a delegate (it was sent from the sender)

			Fanout(InA, OutA, OutB) // Flush the remaining data
			Fanout(InB, OutA, OutB)
			return
		}
	}
}
