package visitor

import "fmt"

type PerimeterCalculator struct {
	perimeter int
}

func (p *PerimeterCalculator) visitForSquare(s *Square) {
	fmt.Println("Calculating perimeter for Square")
}

func (p *PerimeterCalculator) visitForCircle(c *Circle) {
	fmt.Println("Calculating perimeter for Circle")
}

func (p *PerimeterCalculator) visitForRectangle(r *Rectangle) {
	fmt.Println("Calculating perimeter for Rectangle")
}
