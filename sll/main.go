package main

import "fmt"

type Person struct {
	Name string
}

func main() {
	var persons *Person
	change(persons)
	fmt.Println(persons)
	persons
}

func change(person *Person) {
	if person == nil {
		person = &Person{}
		person.Name = "test"
	}
}
