package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":      "Wander",
		"last_name": "Douglas",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	user2 := map[string]map[string]string{
		"name": {
			"first": "Wander",
			"last":  "Santos",
		},
		"course": {
			"name": "Sistemas",
		},
	}

	fmt.Println(user2)
	delete(user2, "name")
	fmt.Println(user2)
}
