package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if args[1] == "" {
		fmt.Println(0)
		return
	}

	words := strings.Split(args[1], " ")

	if args[1] == words[0] {
		fmt.Println(1)
	} else {
		fmt.Println(len(words))
	}

}
