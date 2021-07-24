package state

import "fmt"

type Machine struct {
	current state
}

func NewMachine() *Machine {
	fmt.Println("Machine is ready.")
	return &Machine{newOFF()}
}

func (m *Machine) setCurrent(s state)  {
	m.current = s
}

func (m *Machine) On() {
	m.current.on(m)
}

func (m *Machine) Off() {
	m.current.off(m)
}