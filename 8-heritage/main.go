package main

import "fmt"

// People = people
type People struct {
	Name     string
	LastName string
	Age      uint8
	Height   uint8
}

// Student = student
type Student struct {
	People
	Course string
	Facult string
}

func main() {
	fmt.Println("Heritage")

	people := People{"Wander", "Douglas", 30, 172}
	fmt.Println(people)

	student := Student{people, "Sistema", "ABC"}
	fmt.Println(student)
	fmt.Println(student.Name)

}
