package main

import "fmt"

func main() {
	fmt.Println("Array and Slices")

	var array [5]string
	array[0] = "Position 1"
	fmt.Println(array)

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

	//SLICE
	slice := []int{1, 5, 8, 9, 66}
	fmt.Println(slice)

	slice = append(slice, 47)
	fmt.Println(slice)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	//ARRAY INTERNAL
	fmt.Println("--------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // lenght
	fmt.Println(cap(slice3)) // capacity
}
