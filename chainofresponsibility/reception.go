package chainofresponsibility

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.isRegistered {
		fmt.Println("Patient registration has already done")
		r.next.Execute(p)
		return
	}

	fmt.Println("Reception registering patient")
	p.isRegistered = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}
