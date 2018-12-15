package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"strings"
)

var animals = []string{"lynx", "dog", "cat", "monkey", "dog", "fox", "tiger", "lion"}

func main() {
	fmt.Printf("input: %v\n", animals)
	stream := koazee.StreamOf(animals)
	fmt.Print("stream.Map(strings.Title): ")
	fmt.Println(stream.Map(strings.Title).Do().Out().Val())
}

/**
go run main.go

input: [lynx dog cat monkey dog fox tiger lion]
stream.Map(strings.Title): [Lynx Dog Cat Monkey Dog Fox Tiger Lion]
*/
