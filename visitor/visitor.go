package visitor

type Visitor interface {
	visitForSquare(s *Square)
	visitForCircle(s *Circle)
	visitForRectangle(s *Rectangle)
}
