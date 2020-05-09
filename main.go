package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("===================================================================\n")
	fmt.Printf("Hello %s! This is the Monkey programming language!\n", usr.Username)
	fmt.Printf("Feel free to type in commands\n")
	fmt.Printf("===================================================================\n")
	repl.Start(os.Stdin, os.Stdout)
}
