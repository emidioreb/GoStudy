package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) Show() {
	fmt.Println("This person is the: ", p.name)
	fmt.Println("This person is: ", p.age)
}

func main() {
	j := person{
		name: "joao",
		age:  25,
	}
	j.Show()
}
