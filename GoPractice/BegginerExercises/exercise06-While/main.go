// While
// It will repeat until a condition is true, which could be forever.

package main

import "fmt"

func main() {
	i := 1

	for i < 20 {
		fmt.Println(i)
		i += 1
	}
}
