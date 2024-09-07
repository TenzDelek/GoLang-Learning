package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter your rating")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("the input is %s and its type is : %T", input, input)
	//suppose you want to add 1 but the type of the input are string

	//convert string to int
	//the strings.trimspace is for those who put spaces in the input
	myrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil { //if there is something in the error
		fmt.Println(err)
	} else {
		fmt.Println("the rating is", myrating+1)
	}
}
