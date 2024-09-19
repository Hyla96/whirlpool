package main

import (
	"fmt"
	"os"
	user "os/user"
	"whirlpool/src/repl"
)

func main() {
	u, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Enjoy Whirlpool\n", u.Username)

	repl.Start(os.Stdin, os.Stdout)
}
