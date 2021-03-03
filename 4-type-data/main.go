package main

import (
	"errors"
	"fmt"
)

func main() {
	var number int = 100
	fmt.Println(number)

	var numberUint uint = 58500
	fmt.Println(numberUint)

	//ALIAS
	//int32 = rune
	var number32 rune = 1234
	fmt.Println("int32 = rune", number32)

	//uint8 = byte
	var numberByte byte = 123
	fmt.Println("uint8 = byte", numberByte)

	//REAL NUMBER
	var numberFloat32 float32 = 123.45
	fmt.Println("float32", numberFloat32)

	var numberFloat64 float64 = 1234567.45
	fmt.Println("float64", numberFloat64)

	//STRING
	var str string = "Text"
	fmt.Println(str)

	strTwo := "Text two"
	fmt.Println(strTwo)

	//CHAR = table asci
	char := 'W'
	fmt.Println(char)

	//BOOLEAN
	var bool bool = true
	fmt.Println(bool)

	//ERROR
	var erro error = errors.New("My Error")
	fmt.Println(erro)

}
