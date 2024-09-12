package main

import (
	"fmt"
	"sort"
)

func main() {

	// fmt.Println("hello world")
	//as compare to array we dont have to mention the length of the array
	var fruitlist = []string{"apple", "mango", "banana"}
	// fmt.Printf("the fruit list is %v and its type is %T", fruitlist, fruitlist)

	//adding
	//fruit[0]="mango" //will not work ->this is for array
	fruitlist = append(fruitlist, "amm", "naa")
	fmt.Println(fruitlist)
	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	//implementing make
	highscores := make([]int, 4) //int array with size 4
	highscores[0] = 234
	highscores[1] = 345
	highscores[2] = 455
	highscores[3] = 66
	//default allocation
	// highscores[4]=3333 //results in error as 4 size is allocated
	//but
	highscores = append(highscores, 3434, 5656, 6767)
	//works as append reallocate the memory
	sort.Ints(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))
	fmt.Println(highscores)

}
