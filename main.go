package main

import (
	"fibonacciservice/fibonacci"
	"fmt"
)

func main() {
	r, e := fibonacci.FibonacciSlice(0, 777)
	fmt.Printf("%v %v", r, e)
}
