package main

import "fmt"

type t struct {
	a int
	b string
}

func main() {
	r := t{
		a: 1,
		b: "asdfdas",
	}

	fmt.Println(fmt.Sprintf("%v", r)[1:-1])
}
