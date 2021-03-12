package main

import (
	"fmt"
	"math"
)

type Form interface {
	area() float64
}

type Retangule struct {
	Height float64
	Width  float64
}

func (retangule Retangule) area() float64 {
	return retangule.Width * retangule.Height
}

type Circle struct {
	Raio float64
}

func (circle Circle) area() float64 {
	return math.Pi * (circle.Raio * circle.Raio)
}

func writerArea(form Form) {
	fmt.Printf("Area equal is %0.2f\n", form.area())
}

func main() {
	retangule := Retangule{10, 15}
	writerArea(retangule)

	circle := Circle{10}
	writerArea(circle)
}
