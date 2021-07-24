package observer

type Observer interface {
	GetID() string
	Update(str string)
}
