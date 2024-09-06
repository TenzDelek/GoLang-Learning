package main

import (
	"fmt"
)

//here the first letter should he capital for public so that
//we can use it in the function it same as writing public const token in java

const Tokenexam string = "tenzindelekisthe goat of the house"

func main() {
	fmt.Println(Tokenexam)
	var name string = "tenzin"
	// %T type, %s string ,%t bool, %d int, %f float
	fmt.Printf("hello %s where his type is %T \n", name, name)
	var issmart bool = false
	fmt.Printf("hello %t where his type is %T \n", issmart, issmart)

	var smallints uint8 = 255
	fmt.Printf("hello %d where his type is %T \n", smallints, smallints)

	var floatnum float32 = 12.343433443434343 //we will get only 5 decimal places float64 is more precise
	fmt.Printf("the float number is %f with the type : %T \n", floatnum, floatnum)

	//default value
	var x int
	fmt.Println(x) //always 0 no garbage value
	fmt.Printf("the type of the x is %T \n", x)

	//implicit type : lexer helps
	var username = "tenzin"
	fmt.Println(username)
	username = "tenzindelek" //will work but
	// username=3 //will not work as the type is declare as string by lexer
	fmt.Println(username)

	//no var style
	//:= is for declaration + assignment
	//but only allow to use in function scope outside wont work
	numberofpeople := 10
	fmt.Println(numberofpeople)
}
