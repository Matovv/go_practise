package main

import (
	"fmt"
	"github.com/Matovv/helper-repo-1"
	"math/rand"
	"sort"
)

type Engine struct {
	model string
	power float32
}

func (e *Engine) Start() {
	fmt.Println("Engine",e.model, "started")
} 

func main() {
	fmt.Println("Hello People 😂")
	x := helperRepo1.Multiply(3,10)
	fmt.Println(x,helperRepo1.Version15())

	xi := []string{"lol", "back", "assman"}
	m := map[string]int{
		"civi": 232,
		"bibi": 999,
	}

	for i,v := range xi {
		fmt.Println("index is - ", i, " and value is - ", v)
	}
	for k,v := range m {
		fmt.Println("key is - ", k, " and value is - ", v)
	}

	fmt.Println(rand.Intn(100))

	xiping := []int{5,9,2,12}
	fmt.Println(SortAndReturn(xiping))

	newEngine := Engine{
		model: "Schwarz3",
		power: 75,
	}

	fmt.Println("Engine model -",newEngine.model, "and power -",newEngine.power)
	newEngine.Start();
	
}

func SortAndReturn(x []int) []int {
	a := make([]int,len(x))
	copy(a,x)
	sort.Ints(a)
	return a
} 