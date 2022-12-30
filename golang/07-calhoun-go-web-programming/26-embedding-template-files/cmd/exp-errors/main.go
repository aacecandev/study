package main

import (
	"errors"
	"fmt"
)

func main() {
	err := CreateOrg()
	fmt.Println(err)
}

func Connect() error {
	return errors.New("connection failed")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}
	// ... continue on
	return nil
}

func CreateOrg() error {
	err := CreateUser()
	if err != nil {
		fmt.Printf("integer: %d string: %s any-value: %v\n", 42, "a-string", "another-string")
		return fmt.Errorf("create org: %w", err)
	}
	// ... continue on
	return nil
}
