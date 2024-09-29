package main

import "fmt"

func main() { // the main get invoke auto
	//we cant declate a function inside a function
	fmt.Println("hello there")
	result := adder(4, 5)
	fmt.Println("the answer is", result)
	greeting()
	provalue, message := proadder(1, 2, 3, 4, 5)
	fmt.Println("the sum of the numbers is", provalue)
	fmt.Println("message is", message)
}

func greeting() {
	fmt.Println("tashi delek")
}

func adder(valuea int, valueb int) int { //the output type is necessary if you are returning
	return valuea + valueb
}

func proadder(values ...int) (int, string) { //like props
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "tibet is free"
}
