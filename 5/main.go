package main

import (
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type Person struct {
	ID    int
	First string
	Last  string `json:"last_name"`
	Age   int
	Password []byte
}

// Login function for Person, checking wether input password is the same as the stored hashed password.
func(p Person) Login(password string) string {
	fmt.Printf("User %d attempting to login...\n", p.ID)
	err := bcrypt.CompareHashAndPassword(p.Password, []byte(password))
	if err != nil {
		return fmt.Sprintln("Login failed! Password is incorrect!")
	}
	return fmt.Sprintf("Login Successful! Welcome, %s %s!\n", p.First, p.Last)
}

func (p Person) String() string {
	return fmt.Sprintf("%d - %s %s %d - Password: %v\n", p.ID, p.First, p.Last, p.Age,string(p.Password))
}

func main() {
	person1 := Person{
		ID:    1,
		First: "Jhon",
		Last:  "Snow",
		Age:   32,
		Password: GeneratePassword(1,[]byte("123456")),
	}

	person2 := Person{
		ID:    2,
		First: "Lily",
		Last:  "Swan",
		Age:   19,
		Password: GeneratePassword(1,[]byte("krass32")),
	}

	person3 := Person{
		ID:    3,
		First: "Jackie",
		Last:  "Chan",
		Age:   25,
		Password: GeneratePassword(1,[]byte("Lal#43*Ds")),
	}

	fmt.Println(person1)


	personGroup := []Person{person1, person2, person3}

	for _,person := range personGroup {
		fmt.Println(person.Login("123456"))
	}

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
	JsonUnmarshal(rf, &personGroup2)
	fmt.Println(personGroup2)
}

// Generic Wrapper Function for json.Marshal() with error handling.
func JsonMarshal[T any](obj T) []byte {
	json, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("Error on json marshal:", err)
		return nil
	}
	fmt.Println("Marshal complete!")
	return json
}

// Generic Wrapper Function for json.Unmarshal() with error handling.
func JsonUnmarshal[T any](data []byte, obj *T) {
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("Error on json unmarshal:", err)
	} else {
		fmt.Println("Unmarshal complete!")
	}
}

// Wrapper for bcrypt password generator with error handling.
func GeneratePassword(id int, password []byte) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword(password,10)
	fmt.Printf("Crypting password for User %d ...\n",id)
	if err != nil {
		fmt.Println("Error while crypting password:",err)
		return []byte("")
	}
	fmt.Printf("Password for User %d Set!\n",id)
	return hashedPassword
}


