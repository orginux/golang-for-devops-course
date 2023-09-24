package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("No Args")
		return
	}

	for n, arg := range args[1:] {
		fmt.Println(n, arg)
	}
}
