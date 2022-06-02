package main

import "fmt"

func main(){
	a := 5
	// b := a
	// criando ponteiro
	var b *int
	b = &a
	*b = 10


	// ponteiro muda a informacao que esta no espaço da memoria


	fmt.Println("a = ",a)//10
	fmt.Println("&a = ",&a)//xxx
	fmt.Println("b = ",b)//xxx
	fmt.Println("&b = ",*b)//10
}