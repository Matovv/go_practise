// File create, write, close

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hallow world!")
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer f.Close()

	s := []byte("Hello Legends!")
	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s", err)
	}
}
