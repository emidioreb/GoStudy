// Range
// Range is always used in conjunction with a data structure.

package main

import "fmt"

func main() {

	var arrayNames = [5]string{"Fortaleza", "Ceará", "Santos", "Grêmio", "Santa Cruz"}

	for indice, club := range arrayNames {
		fmt.Println(indice,club)
	}
}
