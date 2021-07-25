package adapter

import "fmt"

type Window struct {}

func (w *Window) insertInSquarePort() {
	fmt.Println("Insert square port into Window machine")
}
