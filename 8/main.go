// Error handling, logging, documentation and other utilities

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Matovv/go_practise/8/example"
)

func main() {
	f, err := os.Create("8/log.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer f.Close()
	log.SetOutput(f)

	Log("Begin Logging ...",nil)

	f2, err := os.Open("noFile.txt")
	if err != nil {
		Log("Error:", err)
		Log("Aborting operation ...",nil)
		FullLogs()
	}
	defer f2.Close()

	fmt.Println(example.ExampleFunc(1,2))

}

// Log function logs the message and error into both the console and log file
func Log(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
		fmt.Println(msg, err)
		return
	}
	log.Println(msg)
	fmt.Println(msg)

}

// FullLogs function logs all the logs from log.txt file into console
func FullLogs() {
	logf, err := os.ReadFile("8/log.txt")
	if err != nil {
		log.Println("Error:", err)
	}
	fmt.Println("Full Logs:")
	fmt.Println(string(logf))
}
