package visitor

import "fmt"

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating area for Square")
}

func (a *AreaCalculator) visitForCircle(c *Circle) {
	fmt.Println("Calculating area for Circle")
}

func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating area for Rectangle")
}
