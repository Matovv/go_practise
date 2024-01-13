package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	ID    int
	First string
	Last  string `json:"last_name"`
	Age   int
}

func main() {

	person1 := Person{
		ID:    1,
		First: "Jhon",
		Last:  "Snow",
		Age:   32,
	}

	person2 := Person{
		ID:    2,
		First: "Lily",
		Last:  "Swan",
		Age:   19,
	}

	person3 := Person{
		ID:    3,
		First: "Jackie",
		Last:  "Chan",
		Age:   25,
	}

	personGroup := []Person{person1, person2, person3}

	fmt.Printf("%+v\n", personGroup)

	f, err := os.Create("5/jsonBlob.txt")
	if err != nil {
		fmt.Println("error:", err)
	}
	defer f.Close()
    
	_, err = f.Write(JsonMarshal(personGroup))
	if err != nil {
		fmt.Println("Error on write:", err)
	} else {
		fmt.Println("Write success!")
	}

	rf, err := os.ReadFile("5/jsonBlob2.txt")
	if err != nil {
		fmt.Println("Error on read:", err)
	} else {
		fmt.Println("Read success!")
		fmt.Println(string(rf))
	}

	var personGroup2 []Person
	JsonUnmarshal(rf,&personGroup2)
	fmt.Println(personGroup2)
}

// Generic Wrapper Function for json.Marshal() with error handling
func JsonMarshal[T any](obj T) []byte {
	json, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Error on json marshal:", err)
		return nil
	}
	fmt.Println("Marshal complete!")
	return json
}

// Generic Wrapper Function for json.Unmarshal() with error handling
func JsonUnmarshal[T any]( data []byte, obj *T) {
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("Error on json unmarshal:", err)
	} else {
		fmt.Println("Unmarshal complete!")
	}
}
