package concurrency

import "fmt"

func main() {

}

func close1() {
	c := make(chan int)
	close(c)

	fmt.Println(<-c) // receive and print??
}
