package main

import "fmt"

func main() {

	// array
	mahasiswa := [3]string{} // with 3 values
	mahasiswa[0] = "komar"
	mahasiswa[1] = "umar"
	mahasiswa[2] = "said"

	// declarative
	// mahasiswa := []string{"arief", "saif", "umar"}

	fmt.Println(mahasiswa)
	fmt.Println(mahasiswa[2])

	// slices
	mhs := [5]int{1, 2, 3, 4, 5}

	fmt.Println(mhs[1:4]) // define slice with column (array output)
}
