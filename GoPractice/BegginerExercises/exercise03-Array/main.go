//Create an array with the number 0 to 10

package main

import "fmt"

func main(){
	var arrayNumber = [4]int{1,2,3}
	arrayNumber[3] = 4

	var arrayNames = [5]string{"Fortaleza", "Ceará", "Santos", "Grêmio", "Santa Cruz" }
	fmt.Println("Number Array: ", arrayNumber)
	fmt.Println("Name Array: ", arrayNames)
}