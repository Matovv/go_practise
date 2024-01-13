// Channels

package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		time.Sleep(time.Second/5)
		c <- 10
	}()

	fmt.Println("received:",<- c)
	
	c2 := make(chan int, 2)
	
	c2 <- 11
	c2 <- 12
	
	fmt.Println("received:",<- c2)
	fmt.Println("received:",<- c2)

}