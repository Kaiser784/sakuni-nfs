package main

import (
	"encoding/base64"
	"log"
	_"crypto/rand"
	"github.com/chzyer/readline"
)

//main user object
type user struct {
	userID uint64				//unique userid
	username string
	accessList map[string]int 	//map containing hash and ACL of files
}

//generate id for user
func generateID() uint64 {
	return uint64(base64.URLEncoding.EncodedLen(64))
}

//create user
func createUser(username string) *user {
	return &user{
		userID: generateID(),
		username: username,
	}
}

//update username
func (currUser *user) updateUser(username string) {
	currUser.username = username
}


func (currentUser * user) initPrompt() (*readline.Instance) {
	autoCompleter := readline.NewPrefixCompleter(
		readline.PcItem("open"),
		readline.PcItem("close"),
		readline.PcItem("mkdir"),
		readline.PcItem("cd"),
		readline.PcItem("rmdir"),
		readline.PcItem("rm"),
		readline.PcItem("quit"),
	)
	prompt, err := readline.NewEx(&readline.Config{
		Prompt: currentUser.username + "[-]$ ",
		HistoryFile: "/tmp/readline.tmp",
		AutoComplete: autoCompleter,
		InterruptPrompt: "^C",
		EOFPrompt: "exit",
	})
	if err != nil {
		log.Fatal(err)
	}
	return prompt
}