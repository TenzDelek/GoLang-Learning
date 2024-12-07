package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	tenz := User{"dummy", "dummy@gmail.com", true, 22}
	fmt.Println(tenz)
	fmt.Printf("the detailed are %+v\n", tenz) //gives the detailed like name:tenz
	fmt.Printf("name is %v and email is %v", tenz.Name, tenz.Email)
	tenz.Greet()
	tenz.Change()
	fmt.Println(tenz.Name) //name is not changed
}

type User struct {
	Name   string //go way is to write in caps to make it public
	Email  string
	Status bool
	Age    int
}

func (u User) Greet() {
	fmt.Println("hello" + u.Name)
}

func (u User) Change() {
	u.Name = "pema"
	fmt.Println(u.Name)
}
