package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid number of arguments.")
	}

	fmt.Println(os.Args)
}

