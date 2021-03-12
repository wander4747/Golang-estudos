package main

import "fmt"

type User struct {
	Name string
	Age  uint8
}

func (user User) save() {
	fmt.Printf("Saving name %s age %d\n", user.Name, user.Age)
}

func (user *User) birtday() {
	user.Age++
}
func main() {
	user := User{"Wander", 30}
	fmt.Println(user)
	user.save()
	user.birtday()
	user.save()
}
