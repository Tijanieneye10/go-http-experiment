package main

import "fmt"

type Author struct {
	Name string
	Age  int
}

func main() {
	author := []Author{
		{
			Name: "Tijani Usman",
			Age:  25,
		},
		{
			Name: "John Doe",
			Age:  29,
		},
	}

	fmt.Println(author)
}
