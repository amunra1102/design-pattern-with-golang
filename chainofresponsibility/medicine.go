package chainofresponsibility

import "fmt"

type Medicine struct {
	next Department
}

func (m *Medicine) Execute(p *Patient) {
	if p.isMedicineProvided {
		fmt.Println("Patient already provided medicine")
		m.next.Execute(p)
		return
	}

	fmt.Println("We are providing medicine to patient")
	p.isMedicineProvided = true
	m.next.Execute(p)
}

func (m *Medicine) SetNext(next Department) {
	m.next = next
}