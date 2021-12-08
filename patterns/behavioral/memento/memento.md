# Шаблон проектирования "Хранитель" в Go

[Оригинал](https://golangbyexample.com/memento-design-pattern-go/)

## Введение

Шаблон "Хранитель" - это поведенческий шаблон проектирования. Он позволяет 
сохранять и восстанавливать предыдущие состояния объектов. В основном он используется
для отмены-повторного проведения каких-либо операций над объектом. Ниже представлены
компоненты, которые используются при создании шаблона проектирования "Хранитель":

* **Создатель**: это реальный объект, состояние которого сохраняется в "Хранитель".
* **Хранитель**: это объект, который хранит состояние создателя.
* **Опекун**: это объект, в котором находятся несколько "Хранителей". По заданному
  индексу он возвращает соответствующего хранителя.
  
В **создателе** определены два методы: `savememento()` и `restorememento()`

* `savememento()` - в этом методе **создатель** сохраняет своё внутреннее состояние в 
  объект "Хранитель".
* `restorememento()` - этот метод принимает в качестве входного параметра объект
  "Хранитель". Создатель, используя переданный "Хранитель", восстанавливает своё 
  предыдущее состояние.
  
# Пример:

**memento/originator.go**

```go
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
```

**memento/memento.go**

```go
type Memento struct {
    state string
}

func (m *Memento) getSavedState() string {
    return m.state
}
```

**memento/caretaker.go**

Обратите внимание, что опекун содержит `mementoArray`, в котором хранятся все 
"Хранители".

```go
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
```

**main.go**

```go
func main() {
    caretaker := memento.NewCaretaker(make([]*memento.Memento, 0))
    originator := memento.NewOriginator("A")
    fmt.Printf("Originator Current State: %s\n", originator.GetState())
  
    caretaker.AddMemento(originator.CreateMemento())
    originator.SetState("B")
  
    fmt.Printf("Originator Current State: %s\n", originator.GetState())
  
    caretaker.AddMemento(originator.CreateMemento())
    originator.SetState("C")
  
    fmt.Printf("Originator Current State: %s\n", originator.GetState())
  
    caretaker.AddMemento(originator.CreateMemento())
  
    originator.RestoreMemento(caretaker.GetMemento(1))
    fmt.Printf("Restored to State: %s\n", originator.GetState())
  
    originator.RestoreMemento(caretaker.GetMemento(0))
    fmt.Printf("Restored to State: %s\n", originator.GetState())
}
```

Результат в терминале:

```shell
go run main.go
Originator Current State: A
Originator Current State: B
Originator Current State: C
Restored to State: B
Restored to State: A
```