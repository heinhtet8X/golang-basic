package main

import "fmt"

func main() {
	x := getBill("Hein")

	name, index := x.getName()

	fmt.Print(x, "\n")
	fmt.Print(index, name, "\n")

	fmt.Println("here two")

	{
		fmt.Println("here one")
		defer fmt.Println("Done...")
	}
}
