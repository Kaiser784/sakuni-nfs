package main

import (
	"fmt"
	"log"
	"github.com/chzyer/readline"
)

func initUser() *user {
	username := setName()
	currUser := createUser(username)
	return currUser
}

//get name from current user
func setName() string {
	var username string

	line, err := readline.New(">")
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("Please enter username: ")
		input, err := line.Readline()
		if err != nil {
			log.Fatal(err)
		}

		if len(input) > 2 {
			fmt.Println("Welcome to sakuni FS ", input)
			username = input
			break
		}
	}
	return username
}