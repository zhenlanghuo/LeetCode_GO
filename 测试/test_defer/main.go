package main

import "fmt"

type person struct {
	name string
}

func main() {
	fmt.Println(f().name)
}

func f() (p person) {
	p = person{name: "11"}

	defer func() {
		p.name = "22"
	}()

	return p
}
