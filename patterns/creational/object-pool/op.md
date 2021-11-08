# Шаблон проектирования "Объектный пул" в Go

[Оригинал](https://golangbyexample.com/golang-object-pool)

Шаблон "Объектный пул" - это порождающий шаблон проектирования, в котором заранее
создаётся набор инициализированных и готовых к использованию объектов. При 
необходимости клиент может запросить объект из пула, использовать его и вернуть
в пул. Объект в пуле никогда не уничтожается.

## Когда стоит использовать

* Когда стоимость создания объекта класса высока, а количество таких объектов,
  которые потребуются в конкретный момент времени, невелико. Примером может быть
  подключение к БД. Создание каждого объекта подключения является дорогостоящим,
  а также может потребоваться не более определенного числа подключений. Шаблон 
  проектирования "Объектный пул" идеально подходит для таких случаев.
* Когда объект в пуле является неизменяемым. Опять в качестве примера рассмотрим
  подключение к БД. Оно является неизменяемым объектом. Практически все его свойства
  не меняются.
* По соображениям производительности. Это значительно повысит производительность 
  приложения, поскольку пул уже создан.
  
## Пример:

**interfaces/poolObject.go**

```go
type PoolObject interface {
    GetID() string // Любой идентификатор, который можно использовать, чтобы отличить двух различных объектов из пула
}
```

**connection/common.go**

```go
type Connection struct {
    Id string
}

func (c *Connection) GetID() string {
    return c.Id
}
```

**connection/pool.go**

```go
type pool struct {
    idle     []interfaces.PoolObject
    active   []interfaces.PoolObject
    capacity int
    mulock   *sync.Mutex
}

// InitPool инициализирует пул
func InitPool(poolObjects []interfaces.PoolObject) (*pool, error) {
    if len(poolObjects) == 0 {
        return nil, errors.New("cannot create a pool of 0 length")
    }
    active := make([]interfaces.PoolObject, 0)
    pool := &pool{
        idle:     poolObjects,
        active:   active,
        capacity: len(poolObjects),
        mulock:   new(sync.Mutex),
    }
    return pool, nil
}

// Loan заимствуем объект из пула
func (p *pool) Loan() (interfaces.PoolObject, error) {
    p.mulock.Lock()
    defer p.mulock.Unlock()
    if len(p.idle) == 0 {
        return nil, errors.New("no pool object free. Please request after sometime")
    }
    obj := p.idle[0]
    p.idle = p.idle[1:]
    p.active = append(p.active, obj)
    fmt.Printf("Loan Pool Object with ID: %s\n", obj.GetID())
    return obj, nil
}

// Retrieve возвращаем объект в пул
func (p *pool) Retrieve(target interfaces.PoolObject) error {
    p.mulock.Lock()
    defer p.mulock.Unlock()
    err := p.Remove(target)
    if err != nil {
        return err
    }
    p.idle = append(p.idle, target)
    fmt.Printf("Return Pool Object with ID: %s\n", target.GetID())
    return nil
}

// Remove удаляем объект из пула
func (p *pool) Remove(target interfaces.PoolObject) error {
    currentActiveLength := len(p.active)
    for i, obj := range p.active {
        if obj.GetID() == target.GetID() {
            p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
            p.active = p.active[:currentActiveLength-1]
            return nil
        }
    }
    return errors.New("target pool object doesn't belong to the pool")
}
```

**main.go**

```go
func main() {
    connections := make([]interfaces.PoolObject, 0)
    for i := 0; i < 3; i++ {
        c := &pool.Connection{Id: strconv.Itoa(i)}
        connections = append(connections, c)
    }
    objectPool, err := pool.InitPool(connections)
    if err != nil {
        log.Fatalf("Init Pool Error: %s", err)
    }
    conn1, err := objectPool.Loan()
    if err != nil {
        log.Fatalf("Pool Loan Error: %s", err)
    }
    conn2, err := objectPool.Loan()
    if err != nil {
        log.Fatalf("Pool Loan Error: %s", err)
    }
    err = objectPool.Retrieve(conn1)
    if err != nil {
        log.Fatalf("Pool Retrieve Error: %s", err)
    }
    err = objectPool.Retrieve(conn2)
    if err != nil {
        log.Fatalf("Pool Retrieve Error: %s", err)
    }
}
```

Результат в терминале:

```shell
go run main.go
Loan Pool Object with ID: 0
Loan Pool Object with ID: 1
Return Pool Object with ID: 0
Return Pool Object with ID: 1
```