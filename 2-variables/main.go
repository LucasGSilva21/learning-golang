package main

import "fmt"

func main() {
	var firstName string = "Lucas"
	lastName := "Silva"
	fmt.Println(firstName)
	fmt.Println(lastName)

	var (
		test1 string = "test1"
		test2 string = "test2"
	)
	fmt.Println(test1)
	fmt.Println(test2)

	test3, test4 := "test3", "test4"
	fmt.Println(test3)
	fmt.Println(test4)

	// inverter os valores
	test3, test4 = test4, test3
	fmt.Println(test3)
	fmt.Println(test4)

	const myConst string = "CONST"
	fmt.Println(myConst)
}
