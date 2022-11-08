package main

import (
	"fmt"
	"os"
)

func main() {
	s, _ := os.Getwd()
	fmt.Println(s)
}
