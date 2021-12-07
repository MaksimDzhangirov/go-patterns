# Шаблон проектирования "Посредник" в Go

[Оригинал](https://golangbyexample.com/mediator-design-pattern-golang/)

## Введение

Шаблон "Посредник" - это поведенческий шаблон проектирования. Он предлагает 
создать объект-посредник для предотвращения прямой связи между объектами, избегая
таким образом прямых зависимостей между ними.

Одним из очень хороших примеров использования шаблона "Посредник" является 
платформа на железнодорожном вокзале. Два поезда никогда напрямую не выясняют
кто из них прибудет на платформу. Диспетчер станции выступает как посредник и
разрешает прибытие на платформу только одному из поездов. Поезд общается с диспетчером
и действует соответствующим образом. Диспетчер также управляет очередью поездов,
ожидающих прибытия на платформу. Как только поезд покидает платформу, он сообщает
одному из поездов, что тот может прибыть на платформу следующим.

Обратите внимание как диспетчер станции (`stationManger`) действует как посредник
между поездами и платформой в приведенном ниже коде.

* пассажирский поезд (`passengerTrain`) и товарный (`goodsTrain`) поезд реализуют
  интерфейс поезда.
* диспетчер станции (`stationManger`) реализует интерфейс посредника.

## Пример:

**interfaces/train.go**

```go
type Train interface {
    RequestArrival()
    Departure()
    PermitArrival()
}
```

**interfaces/mediator.go**

```go
type Mediator interface {
    CanArrival(train Train) bool
    NotifyFree()
}
```

**train/goodsTrain.go**

```go
type goodsTrain struct {
    mediator interfaces.Mediator
}

func NewGoodsTrain(mediator interfaces.Mediator) *goodsTrain {
    return &goodsTrain{
        mediator: mediator,
    }
}

func (t *goodsTrain) RequestArrival() {
    if t.mediator.CanArrival(t) {
        fmt.Println("GoodsTrain: Arrivals")
    } else {
        fmt.Println("GoodsTrain: Waiting")
    }
}

func (t *goodsTrain) Departure() {
    fmt.Println("GoodsTrain: Departures")
    t.mediator.NotifyFree()
}
```

**train/passengerTrain.go**

```go
type passengerTrain struct {
    mediator interfaces.Mediator
}

func NewPassengerTrain(mediator interfaces.Mediator) *passengerTrain {
    return &passengerTrain{
        mediator: mediator,
    }
}

func (t *passengerTrain) RequestArrival() {
    if t.mediator.CanArrival(t) {
        fmt.Println("PassengerTrain: Arrivals")
    } else {
        fmt.Println("PassengerTrain: Waiting")
    }
}

func (t *passengerTrain) Departure() {
    fmt.Println("PassengerTrain: Departures")
    t.mediator.NotifyFree()
}

func (t *passengerTrain) PermitArrival() {
    fmt.Println("PassengerTrain: Arrival Permitted. Arrivals")
}
```

**manager/stationManager.go**

```go
type stationManager struct {
    isPlatformFree bool
    lock           *sync.Mutex
    trainQueue     []interfaces.Train
}

func NewStationManager() *stationManager {
    return &stationManager{
        isPlatformFree: true,
        lock:           &sync.Mutex{},
    }
}

func (s *stationManager) CanArrival(t interfaces.Train) bool {
    s.lock.Lock()
    defer s.lock.Unlock()
    if s.isPlatformFree {
        s.isPlatformFree = false
        return true
    }
    s.trainQueue = append(s.trainQueue, t)
    return false
}

func (s *stationManager) NotifyFree() {
    s.lock.Lock()
    defer s.lock.Unlock()
    if !s.isPlatformFree {
        s.isPlatformFree = true
    }
    if len(s.trainQueue) > 0 {
        firstTrainInQueue := s.trainQueue[0]
        s.trainQueue = s.trainQueue[1:]
        firstTrainInQueue.PermitArrival()
    }
}
```

**main.go**

```go
func main() {
    stationManager := manager.NewStationManager()
    passengerTrain := train.NewPassengerTrain(stationManager)
    goodsTrain := train.NewGoodsTrain(stationManager)
    passengerTrain.RequestArrival()
    goodsTrain.RequestArrival()
    passengerTrain.Departure()
}
```

Результат в терминале:

```shell
go run main.go
PassengerTrain: Arrivals
GoodsTrain: Waiting
PassengerTrain: Departures
GoodsTrain: Arrival Permitted. Arrivals
```