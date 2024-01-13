package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(4)
	x := 0
	go GoTest("1", &x)
	GoTest("2", &x)
	GoTest("3", &x)
	GoTest("4", &x)
	wg.Wait()
	fmt.Println("Final x =", x)

	counter := 0
	const count = 100

	wg.Add(count)

	var mu sync.Mutex

	for i:=0; i<count; i++ {
		go func() {
			mu.Lock()
			v:=counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Final Counter:",counter)


}

func GoTest(s string, x *int) {
	*x++
	fmt.Println("Test -", s)
	wg.Done()
}
