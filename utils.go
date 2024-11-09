package main

import (
	"fmt"
	"log"
)

func GetName() string {
	name := ""

	fmt.Printf("Enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}
