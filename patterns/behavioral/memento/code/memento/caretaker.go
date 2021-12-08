package memento

type caretaker struct {
	mementoArray []*Memento
}

func NewCaretaker(mementoArray []*Memento) *caretaker {
	return &caretaker{
		mementoArray: mementoArray,
	}
}

func (c *caretaker) AddMemento(m *Memento) {
	c.mementoArray = append(c.mementoArray, m)
}

func (c *caretaker) GetMemento(index int) *Memento {
	return c.mementoArray[index]
}
