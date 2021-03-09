package main

import (
	"fmt"
)

func main() {
	/*i := 0

	for i < 10 {
		i++
		fmt.Println("Increment i = ", i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j += 2 {
		fmt.Println("Increment j = ", j)
		time.Sleep(time.Second)
	}*/

	names := [3]string{"Wander", "Erika", "Diego"}

	for index, name := range names {
		fmt.Println(index, name)
	}

	for _, word := range "GOLANG" {
		fmt.Println(word, " = ", string(word))
	}

	user := map[string]string{
		"name":      "Wander",
		"last_name": "Douglas",
	}

	for index, u := range user {
		fmt.Println(index, u)
	}
}
