package main

import "fmt"

func main() {
	// main types
	// - string
	// - bool
	// - int 8/16/32/64 (length)
	// - uint 8/16/32/64 uintptr
	// - float32 64
	// - complex64 128

	// variables key (var/const/:=)
	var name = "Arief"
	var age int64 = 32

	const hello string = "hello world"
	const isstudent bool = true

	addres := "Jepara" // shrothand
	tall := 175.5

	nama, mail, umur, kuliah := "arief", "arief@mail.com", "22", "true"

	fmt.Print(name, tall)
	fmt.Printf("%T", age) // to see a typeof
	fmt.Println(hello, addres)

	fmt.Printf("Hello all, my name is %s with %s, age %s, and iam in collage %s.", nama, mail, umur, kuliah)
}
