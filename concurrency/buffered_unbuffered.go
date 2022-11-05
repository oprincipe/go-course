package concurrency

func main() {

}

func unbufferedFunction() {

	unbuffered := make(chan int)

	// 1) block -> Try to get data from the channel but nobody is sending data
	_ = <-unbuffered

	// 2) blocks -> Try to send data to the channel but nobody is there to collect the data
	unbuffered <- 1

	// 3) it blocks until both functions reach the certain channel operation.
	// 	  Sometimes used to sync 2 goroutines to sync at certain points
	go func() { <-unbuffered }()
	unbuffered <- 1
}

func bufferedFunction() {
	buffered := make(chan int, 1)

	// 4) block cause there's no data
	// a := <-buffered

	// 5) fine cause the buffer wait to receive the data
	buffered <- 1

	// 6) blocks (buffer is full)
	buffered <- 2
}
