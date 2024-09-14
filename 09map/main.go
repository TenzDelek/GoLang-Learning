package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	prices := make(map[string]int)
	prices["tomato"] = 100
	prices["apple"] = 200
	fmt.Println(prices)
	fmt.Println(prices["apple"])
	delete(prices, "apple")
	fmt.Println("after delete: ", prices)

	//loops

	//we can also use _ if we dont need the key
	for key, val := range prices {
		fmt.Println("for key ", key, "the value is", val)
		fmt.Printf("for key %v the value is %v\n", key, val)
	}
}
