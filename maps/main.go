package main

import "fmt"

func main() {
	// map for representative to key and value data, like object in js

	// manual asignment with make() method
	email := make(map[string]string)

	email["arief"] = "arief@mail.com"
	email["joni"] = "joni@mail.com"
	email["amar"] = "amar@mail.com"
	email["edi"] = "edi@mail.com"

	// declarative
	nik := map[string]int{"arief": 232319, "joni": 38228391, "amar": 8377288277, "edi": 3887289}

	fmt.Println(email)
	fmt.Println(email["arief"])
	fmt.Println(nik)
	fmt.Println(len(nik))
	fmt.Println(nik["joni"])

	// delete one
	delete(nik, "amar")

	// add one
	nik["jono"] = 292392918191919

	fmt.Println(nik)
}
