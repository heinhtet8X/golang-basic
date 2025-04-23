package main

import "fmt"

func main() {
	x := getBill("Hein")

	name := x.getName()

	fmt.Print(x, "\n")
	fmt.Print(name, "\n")
}