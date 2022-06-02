// A goroutine is a function that can run concurrently.

package main

import "fmt"

func mensagem(msg string){
	fmt.Println(msg)
}

func main(){
go mensagem("go routine")
mensagem("function")
fmt.Scanln()
}