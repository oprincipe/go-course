package main

//Single comment line
import (
	"fmt"
	"runtime"
)

/*
main Comment with many lines
for a function
*/
func main() {
	fmt.Println(runtime.NumCPU() + 1)
}
