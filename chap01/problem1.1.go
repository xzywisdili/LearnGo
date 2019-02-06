package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Print os.Args[0]
	fmt.Println(os.Args[0])
	fmt.Println(strings.Join(os.Args[1:], " "))
}
