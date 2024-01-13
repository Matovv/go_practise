// Channels, and implementing context on infinite goroutines

package main

import (
	"context"
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

	c3 := make(chan int)
    
	// good logical solution to concurrency with channels
	// send to channel with go routine
	go send(c3)

	// receive without goroutine, this will make code wait until it received content of channel
	// before continuing to exit
	// otherwise the programm will end without possibility for these functions to do their job
    receive(c3)

	c4 := make(chan int)

	go func() {
		for i:=0;i<100;i++ {
			time.Sleep(time.Millisecond*25)
			c4 <- i*5
		}
		close(c4)
	}()

	for v := range c4 {
		fmt.Println("streaming:", v)
	}

	fmt.Println("Finished Streaming! Checking context on infinite goroutine ...")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(time.Millisecond*200)
				fmt.Println("Infinitely printing ...")	
			}
		}	
	}()
		
	time.Sleep(time.Second*2)
	fmt.Println("Canceling infinite printing!")	
	cancel()
	
	time.Sleep(time.Second)
	fmt.Println("Finished! About to exit ...")

}

func send(c chan<- int) {
	c <- 41
}

func receive(c <-chan int) {
	fmt.Println("received:",<-c)
}