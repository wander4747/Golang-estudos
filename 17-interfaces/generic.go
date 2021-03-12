package main

import (
	"fmt"
	"reflect"
)

func main() {
	generic(12)
	generic("String")
	generic(true)
	generic(30.6)
}

func generic(interf interface{}) {
	fmt.Println(interf, reflect.TypeOf(interf))
}
