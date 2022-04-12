package main

import (
	"fmt"
	"github.com/greenteabiscuit/go-interpreter/monkey/repl"
	"os"
	"os/user"
)

func main() {
	curUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Hello %s\n", curUser.Username)
	fmt.Printf("feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
