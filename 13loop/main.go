package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	Day := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	fmt.Println(Day)
	//not much used
	// for i := 0; i < len(Day); i++ {
	// 	fmt.Println(Day[i])
	// }

	// for i := range Day {
	// 	fmt.Println(Day[i])
	// }
	//or
	for index, val := range Day {
		fmt.Printf("the index is %v and the value is %v\n", index, val)
	}

	incrementvalue := 1

	for incrementvalue < 10 { //like a while loop (go doesnt have while loop)
		if incrementvalue == 2 {
			goto lco
		}
		if incrementvalue == 9 {
			break
		}
		if incrementvalue == 6 {
			incrementvalue++
			continue //skip 6
		}
		fmt.Println(incrementvalue)
		incrementvalue++
	}

lco:
	fmt.Println("i got jumped here")
}
