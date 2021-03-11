package main

import "fmt"

func main() {
	fmt.Println(studentIsApproved(6, 6))
}

func recoverRun() {
	if r := recover(); r != nil {
		fmt.Println("recoverRun")
	}
}

func studentIsApproved(n1, n2 float64) bool {
	defer recoverRun()

	average := (n1 + n2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}
	panic("average equals 6")
}
