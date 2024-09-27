package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("if else")
	var res string
	logincount := 24
	if logincount < 10 {
		res = "Regular User"
	} else if logincount > 10 {
		res = "top G"
	} else { //remember that else if and else should be after the { if declare below { not consider by lexar
		res = "admin"
	}

	fmt.Println(res)

	if 9%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	//assign on the go and check
	if num := 3; num < 5 {
		fmt.Println("less than 5")
	} else {
		fmt.Println("greater than 5")
	}

	//nil like python we have None
	var err error = errors.New("i found a error")

	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println("no error")
	}
	errs := "I found an errsor"

	if errs != "" {
		fmt.Println("errsor")
	} else {
		fmt.Println("no errsor")
	}
}
