package state

import "fmt"

type on struct {}

func newOn() state {
	return &on{}
}

func (o *on) on(machine *Machine) {
	fmt.Println("Machine already ON")
}

func (o *on) off(machine *Machine) {
	fmt.Println("Machine is going from ON to OFF")
	machine.setCurrent(newOFF())
}