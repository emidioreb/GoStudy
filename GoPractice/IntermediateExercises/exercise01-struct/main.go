package main

import "fmt"

type person struct{
	nome string
	idade int
	email string
}

func main(){
	joao := person{
		nome: "joao",
		idade:25,
		email: "emidioreb@gmail.com",
	}

	fmt.Println(joao)
}

//2022326915384