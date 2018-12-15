package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
)

type message struct {
	user    string
	message string
}

var messages = []*message{
	{user: "John", message: "Hello Jane"},
	{user: "Jane", message: "Hey John, how are you?"},
	{user: "John", message: "I'm fine! and you?"},
	{user: "Jane", message: "Me too"},
}

func main() {

	stream := koazee.StreamOf(messages)
	stream.ForEach(func(m *message) {
		fmt.Printf("%s: \"%s\"\n", m.user, m.message)
	}).Do()
}

/**
go run main.go

John: "Hello Jane"
Jane: "Hey John, how are you?"
John: "I'm fine! and you?"
Jane: "Me too"
*/
