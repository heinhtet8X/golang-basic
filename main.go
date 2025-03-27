package main

import "fmt"

func main() {

	name := "Hein Htet"

	var address = &name
	var memoName = *address

	name = "Hein Htet Aung"

	fmt.Println("Memo Address", address)
	fmt.Println(memoName)
	fmt.Println(name)

	x := []int{1, 2, 3, 4, 5}

	fmt.Println(x[:3])
	// menu := map[string][string]{
	// 	"1": "One",
	// }

	fmt.Println(x)

	// getting.helloworld("HEIN") // [1 2 3 4 5]
	// helloworld("HEIN") // [1 2 3 4 5]
}
