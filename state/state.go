package state

type state interface {
	on(m *Machine)
	off(m *Machine)
}
