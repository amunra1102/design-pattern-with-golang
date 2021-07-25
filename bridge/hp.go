package bridge

import "fmt"

type HP struct {}

func (h *HP) printFile() {
	fmt.Println("Printing by a HP printer")
}
