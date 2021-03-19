package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangule", func(t *testing.T) {
		retangule := Retangule{10, 12}
		areaResult := retangule.Area()
		areaExpected := float64(120)

		if areaExpected != areaResult {
			t.Fatalf("The area received %f is different from the expected %f ", areaResult, areaExpected)
		}

	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		areaResult := circle.Area()
		areaExpected := float64(math.Pi * 100)

		if areaExpected != areaResult {
			t.Fatalf("The area received %f is different from the expected %f ", areaResult, areaExpected)
		}
	})
}
