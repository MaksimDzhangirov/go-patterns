# Шаблон проектирования "Внедрение Зависимости" в Go

Шаблон "Внедрение Зависимости" - это структурный паттерн проектирования. 
Применяется для реализации слабосвязанной архитектуры, чтобы получить более 
тестируемый, сопровождаемый и расширяемый код.

## Пример:

В структуру `Program` внедряются `Db` и `Greeter`. Поскольку для полей
используются интерфейсы, при инициализации туда могут быть переданы различные
реализации, в том числе и тестовые. Это позволяет легко расширять и тестировать
приложение.

**interfaces/db.go**

```go
type Database interface {
    GetUser() string
}
```

**interfaces/greeter.go**

```go
type Greeter interface {
    Greet(userName string)
}
```

**db/default.go**

```go
type defaultDatabase struct {
}

func NewDefaultDatabase() *defaultDatabase {
    return &defaultDatabase{}
}

func (db *defaultDatabase) GetUser() string {
    return "user 1"
}
```

**greeter/nice_greeter.go**

```go
type niceGreeter struct {
}

func NewNiceGreeter() *niceGreeter {
    return &niceGreeter{}
}

func (ng *niceGreeter) Greet(userName string) {
    log.Printf("Hello %s! Nice to meet you", userName)
}
```

**main.go**

```go
func main() {
    ddb := db.NewDefaultDatabase()
    nGreeter := greeter.NewNiceGreeter()
    prog := program.NewProgram(ddb, nGreeter)
    
    prog.Execute()
}
```

Результат в терминале:

```shell
go run main.go
2021/11/09 19:17:53 Hello user 1! Nice to meet you
```