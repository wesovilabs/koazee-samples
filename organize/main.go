package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"strings"
)

var animals = []string{"lynx", "dog", "cat", "monkey", "fox", "tiger", "lion"}

func main() {
	fmt.Print("input: ")
	fmt.Println(animals)
	stream := koazee.StreamOf(animals)

	fmt.Print("stream.Reverse(): ")
	fmt.Println(stream.Reverse().Out().Val())

	fmt.Print("stream.Sort(strings.Compare): ")
	fmt.Println(stream.Sort(strings.Compare).Out().Val())

}

/**
go run main.go

input: [lynx dog cat monkey fox tiger lion]
stream.Reverse(): [lion tiger fox monkey cat dog lynx]
stream.Sort(strings.Compare): [cat dog fox lion lynx monkey tiger]
*/
