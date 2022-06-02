// Check if a file exists on your local disk

package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("./db.json"); err == nil {
		fmt.Println("arquivo encontrado")
	} else {
		fmt.Println("arquivo nao encontrado")
	}
}
