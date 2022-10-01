package main

import "fmt"

type Person struct {
	name string
}

func (n *Person) set(t string) {
	n.name = t
}
func (n *Person) HandleEvent(vacancies []string) {
	fmt.Println("Hi dear " + n.name)
	fmt.Println("Vacancies updated: ")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
}
