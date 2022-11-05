package main

import "fmt"

func main() {
	//  conditions (if/switch)
	x := 5
	y := 10

	if x < y {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// else if
	color := "blue"
	if color == "blue" {
		fmt.Println(true)
	} else if color == "red" || color == "green" {
		fmt.Println("not red or green")
	} else {
		fmt.Println(false)
	}

	// switch case
	switch x {
	case 1:
		fmt.Println("nope")
	case 2:
		fmt.Println("nope")
	case 5:
		println("oke")
	default:
		print("nope")
	}
}
