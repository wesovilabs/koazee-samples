package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 5, 4, 3, 2, 7, 1, 8, 2, 3}

func main() {
	fmt.Printf("input: %v\n", numbers)

	stream := koazee.StreamOf(numbers)
	fmt.Print("stream.Add(10): ")
	fmt.Println(stream.Add(10).Do().Out().Val())

	fmt.Print("stream.Drop(5): ")
	fmt.Println(stream.Drop(5).Do().Out().Val())

	fmt.Print("stream.DeleteAt(4): ")
	fmt.Println(stream.DeleteAt(4).Do().Out().Val())

	fmt.Print("stream.Set(0,5): ")
	fmt.Println(stream.Set(0, 5).Do().Out().Val())

	fmt.Print("stream.Pop(): ")
	val, newStream := stream.Pop()
	fmt.Printf("%d ... ", val.Int())
	fmt.Println(newStream.Out().Val())

}

/**
go run main.go

input: [1 5 4 3 2 7 1 8 2 3]
stream.Add(10): [1 5 4 3 2 7 1 8 2 3 10]
stream.Drop(5): [1 4 3 2 7 1 8 2 3]
stream.DeleteAt(4): [1 5 4 3 7 1 8 2 3]
stream.Set(0,5): [5 5 4 3 2 7 1 8 2 3]
stream.Pop(): 5 ... [5 4 3 2 7 1 8 2 3]
*/
