package visitor

type Shape interface {
	Accept(v Visitor)
}
