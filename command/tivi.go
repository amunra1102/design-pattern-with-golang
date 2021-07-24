package command

import "fmt"

type Tivi struct {
	isRunning bool
}

func (t *Tivi) on() {
	t.isRunning = true
	fmt.Println("Turing tivi on")
}

func (t *Tivi) off() {
	t.isRunning = false
	fmt.Println("Turing tivi off")
}
