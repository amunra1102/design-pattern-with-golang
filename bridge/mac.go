package bridge

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for Mac")
	m.printer.printFile()
}

func (m *Mac) SetPrinter(printer Printer) {
	m.printer = printer
}