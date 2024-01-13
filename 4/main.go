// Database simulation, more complex interface


package main

import (
	"fmt"
	"log"
)

// User represents an user with an id, first name, last name and age
type User struct {
	ID    int
	First string
	Last  string
	Age   int
}

// MockDatabase is a temporary service that stores retrievable data.
type MockDatabase struct {
	Users map[int]User
}

func (md MockDatabase) GetUser(id int) (User, error) {
	fmt.Printf("Getting user %v ...\n", id)
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	fmt.Printf("User %v is found!\n", id)
	return user, nil
}

func (md MockDatabase) UpdateUser(u User) (User, error) {
	_, ok := md.Users[u.ID]
	if ok {
		fmt.Printf("User %v found! Updating ...\n", u.ID)
	} else {
		fmt.Printf("User %v not found! Creating ...\n", u.ID)
	}

	md.Users[u.ID] = u
	user := md.Users[u.ID]

	return user, nil
}

func (md MockDatabase) DeleteUser(id int) error {
	_, ok := md.Users[id]
	if !ok {
		return fmt.Errorf("User %v not found", id)
	}

	delete(md.Users, id)
	fmt.Printf("User %v have been successfully deleted!\n", id)
	return nil
}

// Database defines an interface for storing retrievable data.
// Any TYPE that implements this interface (has these 3 methods) is also of TYPE 'Database'
// So MockDatabase is also of TYPE 'Database'
// If other databases exist besided MockDatabase, they still will be considered of TYPE 'Database'
// and will also have these methods, as long as they too implement these methods
// with this interface we can put any database into this function and
type Database interface {
	GetUser(id int) (User, error)
	UpdateUser(u User) (User, error)
	DeleteUser(id int) error
}

// Service represents a service that manipulates retrievable data from database.
// we can attach any database to it.
// think of repository
type Service struct {
	db Database
}

func (s Service) GetUser(id int) (User, error) {
	return s.db.GetUser(id)
}

func (s Service) UpdateUser(u User) (User, error) {
	return s.db.UpdateUser(u)
}

func (s Service) DeleteUser(id int) error {
	return s.db.DeleteUser(id)
}

func main() {
	mdb := MockDatabase{
		Users: make(map[int]User),
	}

	srvc := Service{
		db: mdb,
	}

	u1 := User{
		ID:    0,
		First: "Jhon",
		Last:  "Phantom",
		Age:   28,
	}

	_, err := srvc.UpdateUser(u1)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	u1Returned, err := srvc.GetUser(0)
	if err != nil {
		log.Fatalf("error %s", err)
	}

	fmt.Println(u1)
	fmt.Println(u1Returned)
	
}
