package chainofresponsibility

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.isPaid {
		fmt.Println("Patient already paid their bill")
		c.next.Execute(p)
		return
	}

	fmt.Println("Patient is paying the bill")
	p.isPaid = true
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
