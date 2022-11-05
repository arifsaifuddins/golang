package main

import "fmt"

func main() {
	// pointer is memory sign number
	x := 5
	y := &x

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*y) //print value for pointer with *
	fmt.Println(*&x)

	// assign and to change x value
	*y = 20

	fmt.Println(x)
}
