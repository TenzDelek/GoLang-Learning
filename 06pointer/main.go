package main

import "fmt"

func main() {
	// var b *int //creating a pointer
	// fmt.Println(b)

	val := 10
	var ref = &val                     //the & is call referening
	fmt.Println("the pointer is", ref) //memory address
	fmt.Println("the value is", *ref)

	*ref = *ref + 2                  //the orgi value is being change
	fmt.Println("the value is", val) //our orgi value changes
}
