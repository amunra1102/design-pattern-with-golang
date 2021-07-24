package prototype

type INode interface {
	Clone() INode
	Print(str string)
}
