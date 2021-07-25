package bridge

type Computer interface {
	print()
	setPrinter(printer Printer)
}
