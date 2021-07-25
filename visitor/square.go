package visitor

type Square struct {
	Side int
}

func (s *Square) Accept(v Visitor) {
	v.visitForSquare(s)
}