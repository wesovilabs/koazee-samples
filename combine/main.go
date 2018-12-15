package main

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"strings"
)

type Person struct {
	Name string
	Male bool
	Age  int
}

var people = []*Person{
	{"John Smith", true, 32},
	{"Peter Pan", true, 17},
	{"Jane Doe", false, 20},
	{"Anna Wallace", false, 35},
	{"Tim O'Brian", true, 13},
	{"Celia Hills", false, 15},
}

func main() {
	stream := koazee.
		StreamOf(people).
		Filter(func(person *Person) bool {
			return !person.Male
		}).
		Sort(func(person, otherPerson *Person) int {
			return strings.Compare(person.Name, otherPerson.Name)
		}).
		ForEach(func(person *Person) {
			fmt.Printf("%s is %d years old\n", person.Name, person.Age)
		})

	fmt.Println("Operations are not evaluated until we perform stream.Do()\n")
	stream.Do()
}

/**
go run main.go

Operations are not evaluated until we perform stream.Do()

Anna Wallace is 35 years old
Celia Hills is 15 years old
Jane Doe is 20 years old
 */