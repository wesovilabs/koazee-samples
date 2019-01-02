package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	fmt.Printf("input: %v\n", numbers)
	stream := koazee.StreamOf(numbers)
	count, _ := stream.Count()
	fmt.Printf("stream.Count(): %d\n", count)
	index, _ := stream.IndexOf(2)
	fmt.Printf("stream.IndexOf(2): %d\n", index)
	indexex, _ := stream.IndexesOf(2)
	fmt.Printf("stream.IndexesOf(2): %d\n", indexex)
	index, _ = stream.LastIndexOf(2)
	fmt.Printf("stream.LastIndexOf(2): %d\n", index)
	contains, _ := stream.Contains(7)
	fmt.Printf("stream.Contains(7): %v\n", contains)
}

/**
go run main.go

input: [1 5 4 3 2 7 1 8 2 3]
stream.Count(): 10
stream.IndexOf(2): 4
stream.IndexesOf(2): [4 8]
stream.LastIndexOf(2): 8
stream.Contains(7): true
*/
