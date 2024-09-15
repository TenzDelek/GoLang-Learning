package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	tenzin := User{"tenzin", "tenzindelek@gmail.com", true, 22}
	fmt.Println(tenzin)
	fmt.Printf("the detailed are %+v\n", tenzin) //gives the detailed like name:tenzin
	fmt.Printf("name is %v and email is %v", tenzin.Name, tenzin.Email)
}

type User struct {
	Name   string //go way is to write in caps to make it public
	Email  string
	Status bool
	Age    int
}
