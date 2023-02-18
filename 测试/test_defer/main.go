package main

import "fmt"

type person struct {
	name string
}

func main() {
	fmt.Println(f().name)
}

func f() *person {
	p := person{name: "11"}

	defer func() {
		fmt.Println(p)
		p.name = "22"
	}()

	return &p
}
