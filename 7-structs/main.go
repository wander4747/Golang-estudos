package main

import "fmt"

// User = user
type User struct {
	Name    string
	Age     uint8
	Address Address
}

// Address = address
type Address struct {
	Street string
	Number uint8
}

func main() {
	user := User{
		Name: "Wander",
		Age:  30,
	}
	fmt.Println(user)

	//method 2
	address := Address{"Rua ABC", 123}
	n := User{"Wander", 30, address}
	fmt.Println(n)

	//method 3
	var u User
	u.Name = "Wander"
	u.Age = 30

	fmt.Println(u)

	//method 4
	m := User{Name: "Wander"}
	fmt.Println(m)

}
