// Read file

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	readFile, err := ioutil.ReadFile("./db.json")

	if err != nil{
		println(err)
	}

	str := string(readFile)

	fmt.Println(str)

}
