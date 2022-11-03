package main

import "fmt"

func main() {
	number := 100

	for no := 1; no < number; no++ {
		fmt.Println(no)
	}

	nom := 98
	for nom < number {
		fmt.Println(nom)
		nom++
	}
}
