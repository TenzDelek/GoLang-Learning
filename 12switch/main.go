package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	switch dicenumber {
	case 1:
		fmt.Println("you got 1")
	case 2:
		fmt.Println("you got 2")
	case 3:
		fmt.Println("you got 3")
	case 4:
		fmt.Println("you got 4")
	case 5:
		fmt.Println("you got 5")
		fallthrough //print case 6 as well
	case 6:
		fmt.Println("you got 6")
	default:
		fmt.Println("what was that?")
	}

}
