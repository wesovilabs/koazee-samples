package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"strings"
)

var animals = []string{"lynx", "dog", "cat", "monkey", "dog", "fox", "tiger", "lion"}

type Person struct {
	firstName string
	age int
}

func main() {
	fmt.Printf("input: %v\n", animals)
	stream := koazee.StreamOf(animals)
	fmt.Print("stream.Map(strings.Title): ")
	fmt.Println(stream.Map(strings.Title).Do().Out().Val())
	fmt.Print("stream.GroupBy(strings.Len): ")
	out, _ := stream.GroupBy(func(val string)int{return len(val)})
	fmt.Println(out)
}

/**
go run main.go

input: [lynx dog cat monkey dog fox tiger lion]
stream.Map(strings.Title): [Lynx Dog Cat Monkey Dog Fox Tiger Lion]
stream.GroupBy(strings.Len): map[5:[tiger] 4:[lynx lion] 3:[dog cat dog fox] 6:[monkey]]
*/
