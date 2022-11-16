package main

import "fmt"

type t struct {
	a int
	s string
}

func main() {
	r := t{
		a: 1,
	}
	fmt.Println(r)
}
