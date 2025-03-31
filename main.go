package main

import "fmt"

func updateVal(x *string) {
	*x = "Aung"
}

func main() {
	name := "Hein Htet"

	memo_address := &name

	fmt.Println("Current :", name)

	updateVal(memo_address)

	fmt.Println("Updated :", name)
}
