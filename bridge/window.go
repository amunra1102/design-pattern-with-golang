package bridge

import "fmt"

type Window struct {
	printer Printer
}

func (w *Window) Print() {
	fmt.Println("Print request for windows")
	w.printer.printFile()
}

func (w *Window) SetPrint(printer Printer) {
	w.printer = printer
}