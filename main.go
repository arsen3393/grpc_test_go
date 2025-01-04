package main

import "fmt"

type User struct {
	Name string
	Age  int32
}

func main() {
	u := &User{
		Age:  32,
		Name: "Arsen",
	}
	fmt.Println(u)
}
