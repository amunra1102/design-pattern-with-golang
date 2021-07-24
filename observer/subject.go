package observer

type Subject interface {
	Register(observer Observer)
	DeRegister(observer Observer)
	NotifyAll()
}

