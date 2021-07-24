package memento

type Originator struct {
	state string
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{
		state: o.state,
	}
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.state = m.state
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	o.state = state
}
