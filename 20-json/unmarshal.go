package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"-"`
	Age   uint   `json:"age"`
}

func main() {
	dogJSON := `{"name":"Benio","breed":"Vira lata","age":2}`

	var dog Dog

	if err := json.Unmarshal([]byte(dogJSON), &dog); err != nil {
		log.Fatal(err)
	}

	fmt.Printf(dogJSON)

	dogTwoJSON := `{"breed":"Pug","name":"Toby"}`
	dogTwo := make(map[string]string)

	if err := json.Unmarshal([]byte(dogTwoJSON), &dogTwo); err != nil {
		log.Fatal(err)
	}

	fmt.Printf(dogTwoJSON)

}
