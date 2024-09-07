package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// name := `
	// tenzin delek`
	// //with backtick you can write in multiple lines
	// //without backtick (normal)  error
	// fmt.Print(name)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the price for this jordan: ")

	//comma ok | comma error
	//go doesnt have trycatch it does this by using this
	//here the _ is used to ignore the error but we can write error and use that
	price, _ := reader.ReadString('\n') //says read until \n
	fmt.Printf("the price is: %T", price)
}
