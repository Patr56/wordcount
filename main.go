package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args
	words := strings.Split(args[1], " ")

	fmt.Println(len(words))
}
