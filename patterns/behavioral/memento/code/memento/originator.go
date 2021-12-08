package memento

type originator struct {
	state string
}

func NewOriginator(state string) *originator {
	return &originator{
		state: state,
	}
}

func (o *originator) CreateMemento() *Memento {
	return &Memento{state: o.state}
}

func (o *originator) RestoreMemento(m *Memento) {
	o.state = m.getSavedState()
}

func (o *originator) SetState(state string) {
	o.state = state
}

func (o *originator) GetState() string {
	return o.state
}
