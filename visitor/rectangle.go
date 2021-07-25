package visitor

type Rectangle struct {
	A int
	B int
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.visitForRectangle(r)
}