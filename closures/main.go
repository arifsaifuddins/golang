package main

import "fmt"

// returning func in func
func Closures() func(int) int {
	sum := 3
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// init a closure
	sum := Closures()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

	fmt.Println(sum(2))
}
