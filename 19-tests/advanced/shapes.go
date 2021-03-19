package shapes

import (
	"fmt"
	"math"
)

type Form interface {
	Area() float64
}

type Retangule struct {
	Height float64
	Width  float64
}

func (retangule Retangule) Area() float64 {
	return retangule.Width * retangule.Height
}

type Circle struct {
	Raio float64
}

func (circle Circle) Area() float64 {
	return math.Pi * (circle.Raio * circle.Raio)
}

func WriterArea(form Form) {
	fmt.Printf("Area equal is %0.2f\n", form.Area())
}

func main() {
	retangule := Retangule{10, 15}
	WriterArea(retangule)

	circle := Circle{10}
	WriterArea(circle)
}
