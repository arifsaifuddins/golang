package main

import "fmt"

func helloAll(name string, age int) string {
	return "Hello world!, iam " + name + " and iam " + fmt.Sprint(age)
}

func sum(one int, two int) int {
	return one + two
}

func main() {
	fmt.Println(helloAll("Arief Saifuddien", 22))
	fmt.Println(sum(2, 4))
}
