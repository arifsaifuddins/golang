package main

import (
	"fmt"
	"strconv"
)

// sturct is like classes in another language

// define struct with type key and struct
type Student struct {

	// // properties
	// firstname string
	// lastname  string
	// nickname  string
	// city      string
	// age       int
	// isFree    bool

	// if same type value
	firstname, lastname, nickname, city string
	age                                 int
	isFree                              bool
}

// create method (value recieved)
func (s Student) greating() string {
	return "Hi All, iam " + s.firstname + " iam from " + s.city + " and iam " + strconv.Itoa(s.age)
}

// create method (pointer recieved) change
func (d *Student) changeAge() {
	d.age++
}

func main() {
	// declare a struct
	student1 := Student{"arief", "saifuddien", "rif", "jepara", 22, true}
	student2 := Student{"ahmad", "udin", "mat", "khartoum", 21, false}

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(student2.city)

	student1.changeAge()
	student2.changeAge()
	student2.changeAge()
	student2.changeAge()
	fmt.Println(student2.age)

	// print method
	fmt.Println(student1.greating())
	fmt.Println(student2.greating())

	// anonymous struct
	datadiri := struct {
		Name string
		Age  int
	}{
		"arief saifuddien",
		22,
	}

	fmt.Println(datadiri)
}
