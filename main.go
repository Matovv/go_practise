package main

import (
	"fmt"
	"github.com/Matovv/helper-repo-1"
	"math/rand"
)

func main() {
	fmt.Println("Hello People 😂")
	x := helperRepo1.Multiply(3,10)
	fmt.Println(x,helperRepo1.Version15())

	xi := []string{"lol", "back", "assman"}
	m := map[string]int{
		"civi": 232,
		"cumshet": 999,
	}

	for i,v := range xi {
		fmt.Println("index is - ", i, " and value is - ", v)
	}
	for k,v := range m {
		fmt.Println("key is - ", k, " and value is - ", v)
	}

	fmt.Println(rand.Intn(100))
	
}