package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	stream := koazee.StreamOf(numbers)
	fmt.Printf("stream.At(4): %d\n", stream.At(4).Int())
	fmt.Printf("stream.First: %d\n", stream.First().Int())
	fmt.Printf("stream.Last: %d\n", stream.Last().Int())
}

/**
go run main.go

stream.At(4): 2
stream.First: 1
stream.Last: 3
*/