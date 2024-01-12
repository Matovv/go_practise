package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"

	"github.com/Matovv/helper-repo-1"
)

type Engine struct {
	model string
	power float32
}

func (e *Engine) Start() {
	fmt.Println("Engine",e.model, "started")
} 

type MyType int

func (mt *MyType) Multi() {
	*mt*=2
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

	someVar := 10.0;
	someVar2 := &someVar
    fmt.Println("Original before -",someVar)
    fmt.Println("Modified -",Multiply(&someVar))
    fmt.Println("Original after -",someVar)
    fmt.Println("SomeVar2 before -",*someVar2)
	*someVar2 *= 2
    fmt.Println("SomeVar2 after -",*someVar2)
    fmt.Println("SomeVar1 after -",someVar)


	var mt MyType = 5;
	log.Println(mt)
	mt.Multi()
	log.Println(mt)
	mt.Multi()
	log.Println(mt)
	
}

func Multiply(x *float64) float64 {
	*x *= 5
	return *x
}

func SortAndReturn(x []int) []int {
	a := make([]int,len(x))
	copy(a,x)
	sort.Ints(a)
	return a
} 