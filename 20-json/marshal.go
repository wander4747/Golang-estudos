package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	dog := Dog{"Benio", "Vira lata", 2}
	fmt.Println(dog)

	dogJSON, err := json.Marshal(dog)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(dogJSON))
	fmt.Println(bytes.NewBuffer(dogJSON))

	dogTwo := map[string]string{
		"name":  "Toby",
		"breed": "Pug",
	}

	dogTwoJSON, err := json.Marshal(dogTwo)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(dogTwoJSON))
}
