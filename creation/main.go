package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	fmt.Printf("slice: %v\n", numbers)
	stream := koazee.StreamOf(numbers)
	fmt.Printf("stream: %v\n", stream.Out().Val())
}

/**
go run main.go

slice: [1 5 4 3 2 7 1 8 2 3]
stream: [1 5 4 3 2 7 1 8 2 3]
*/
