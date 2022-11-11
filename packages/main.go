package main

// import "fmt" // is ith one package only

// import package from go package
import (
	"fmt"
	"math"
	"math/rand"

	"github.com/ariefsaifuddien/golang/packages/arief" // must init go mod init
)

func main() {
	fmt.Println(math.Floor(2.33))
	fmt.Println(math.Ceil(2.33))
	fmt.Println(rand.Int63())

	fmt.Println(arief.Hellocall())
}
