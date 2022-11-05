package main

import "fmt"

func main() {
	number := 10

	for no := 1; no < number; no++ {
		fmt.Println(no)
	}

	nom := 98
	for nom < number {
		fmt.Println(nom)
		nom++
	}

	// fizzbuzz
	for i := 1; i < 50; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

	// range for

	// slice
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	for i, v := range nums {
		fmt.Printf("%d : %d\n", i, v)
	}

	// without print an index
	for _, v := range nums {
		fmt.Printf("no index for : %d\n", v)
	}

	// maps
	names := map[string]string{"jamal": "jamaluddin", "saif": "saifuddien", "edi": "edi wibowo"}

	for k, val := range names {
		fmt.Printf("%s key for %s\n", k, val)
	}
}
