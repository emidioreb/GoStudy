package main

import "fmt"

func numberSlice() {
	myset := []int{0, 1, 2, 3, 4}

	x := myset[0:2]
	fmt.Println(x)
}

func textSlice() {
	myText := "Hello word"
	fmt.Println(len(myText))

	hello := myText[0:5]
	word := myText[6:10]

	fmt.Println(hello)
	fmt.Println(word)
}

func main() {
	numberSlice()
	textSlice()
}
