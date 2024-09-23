package main

import (
	"fmt"
	"github.com/Hyla96/whirlpool/repl"
	"os"
	user "os/user"
)

func main() {
	u, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Enjoy Whirlpool\n", u.Username)

	repl.Start(os.Stdin, os.Stdout)
}
