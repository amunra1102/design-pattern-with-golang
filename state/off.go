package state

import "fmt"

type off struct {}

func newOFF() state {
	return &off{}
}

func (o *off) on(machine *Machine) {
	fmt.Println("Machine is going from OFF to ON")
	machine.setCurrent(newOn())
}

func (o *off) off(machine *Machine) {
	fmt.Println("Machine already OFF")
}