package chainofresponsibility

import "fmt"

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.isDoctorChecked {
		fmt.Println("Patient already checked by doctor")
		d.next.Execute(p)
		return
	}

	fmt.Println("Doctor is checking patient")
	p.isDoctorChecked = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}
