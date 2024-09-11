package main

import "fmt"

//not used much in golang array
func main() {
	fmt.Println("Hello World")
	var food [4]string //in go it is neccesary to declare the size
	food[0] = "pizza"
	food[1] = "burger"
	//if we dont put the 2nd place value it will leave a space when printing
	//the array.[pizza burger  noodles]
	food[3] = "noodles"
	fmt.Println(food)
	fmt.Println(food[1])
	fmt.Println(len(food)) //but the length will be still 4 as we have declare it as 4

	var vegtable = [3]string{"carrot", "potato", "onion"}
	fmt.Println(vegtable)
}
