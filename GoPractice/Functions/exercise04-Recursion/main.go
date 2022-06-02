package main

import "fmt"

func countdown(x int) int {
	if x == 0 {
		return 0
	}
	fmt.Println(x)
	return countdown(x - 1)
}

func main() {
	countdown(10)
}
