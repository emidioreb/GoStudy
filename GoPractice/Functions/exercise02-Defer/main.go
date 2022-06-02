package main

import (
	"fmt"
	"os"
)

func hello() {
	fmt.Println("Hello")
}

func helloComDefer() {
	fmt.Println("Hello com defer")
}

func create() {
	f, _ := os.Create("hello.txt")
	defer f.Close()
	fmt.Fprintln(f, "hello world")
}

func main() {
	defer helloComDefer()
	hello()
	create()
}

