package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var animals = []string{"lynx", "dog", "cat", "monkey", "dog", "fox", "tiger", "lion"}

func main() {
	fmt.Print("input: ")
	fmt.Println(animals)
	stream := koazee.StreamOf(animals)

	fmt.Print("stream.Take(1,4): ")
	fmt.Println(stream.Take(1, 4).Out().Val())

	fmt.Print("stream.Filter(len==4): ")
	fmt.Println(stream.
		Filter(
			func(val string) bool {
				return len(val) == 4
			}).
		Out().Val(),
	)
	fmt.Print("stream.RemoveDuplicates(): ")
	fmt.Println(stream.RemoveDuplicates().Out().Val())
}

/**
go run main.go

input: [lynx dog cat monkey dog fox tiger lion]
stream.Take(1,4): [dog cat monkey dog]
stream.Filter(len==4): [lynx lion]
stream.RemoveDuplicates(): [lynx dog cat monkey fox tiger lion]
*/
