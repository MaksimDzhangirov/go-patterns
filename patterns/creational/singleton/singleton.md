# Шаблон проектирования "Одиночка" в Go

Шаблон "Одиночка" - это порождающий шаблон проектирования, а также один из наиболее
часто используемых шаблонов. Этот шаблон используется, когда должен существовать
только один экземпляр структуры. Этот единственный экземпляр называется 
объектом-одиночкой. Некоторые случаи, когда применим объект-одиночка:

1. **Экземпляр БД** — мы хотим создать только один экземпляр объекта БД и этот экземпляр 
   будет использоваться во всём приложении.
2. **Экземпляр Logger** - опять же следует создавать только один экземпляр объекта
   Logger и он должен использоваться во всём приложении.

Объект-одиночка создаётся при первой инициализации структуры. Обычно в структуре,
для которого необходимо создать только один экземпляр, определен метод 
`getInstance()`. При вызове `getInstance()` в первый раз создаёт, а дальше просто
возвращает уже созданный один и тот же экземпляр объекта-одиночки.

В Go существуют горутины. Следовательно, структура, создающая объект-одиночку,
должна возвращать один и тот же экземпляр каждый раз, когда несколько горутин 
пытаются получить доступ к ней. При этом очень легко допустить ошибку в коде.
Приведенный ниже фрагмент кода показывает как правильно создать объект-одиночку:

```go
var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating Single Instance Now")
            singleInstance = &single{}
        } else {
            fmt.Println("Single Instance already created-1")
        }
    } else {
        fmt.Println("Single Instance already created-2")
    }
    return singleInstance
}
```

Вышеприведенный код гарантирует, что будет создан только один экземпляр 
структуры. Отметим следующие моменты:

1. В начале идёт проверка на `nil` `singleInstance`. Это сделано для предотвращения
   дорогостоящих операций блокировки при каждом вызове метода `getInstance()`. 
   Если условие не выполняется, это означает, что `singleInstance` уже создан.
2. `singleInstance` создаётся после блокировки мутекса.
3. После блокировки выполняется еще одна проверка на `nil` `singleInstance`. Это 
   гарантирует, что даже если более одной горутины пройдёт первую проверку,
   то только одна горутина сможет создать экземпляр одиночки. В противном случае
   каждая из горутин создаст свой собственный экземпляр.
   
Ниже приведена полная версия:

**single.go**

```go
package main

import (
    "fmt"
    "sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {
    if singleInstance == nil {
        lock.Lock()
        defer lock.Unlock()
        if singleInstance == nil {
            fmt.Println("Creating Single Instance Now")
            singleInstance = &single{}
        } else {
            fmt.Println("Single Instance already created-1")
        }
    } else {
        fmt.Println("Single Instance already created-2")
    }
    return singleInstance
}
```

**main.go**

```go
package main

import "fmt"

func main() {
    for i := 0; i < 100; i++ {
        go getInstance()
    }
    // Scanln аналогичен Scan, но останавливает сканирование, когда находит символ новой строки,
    // Таким образом, программа будет ожидать ввода символа новой строки и только после этого завершит работу
    fmt.Scanln()
}
```

Результат в терминале:
```shell
go run .
Creating Single Instance Now
Single Instance already created-2
Single Instance already created-2
Single Instance already created-2
Single Instance already created-2
Single Instance already created-2
Single Instance already created-2
Single Instance already created-2
...
```

Обратите внимание:

1. Программа выводит "Creating Single Instance Now", что означает, что только 
   одна горутина смогла создать экземпляр объекта-одиночки.
2. Наличие одной или нескольких строк "Single Instance already created-1" означает, 
   что для одной или несколько горутин значение singleInstance было равно nil при первой проверке.
3. Наличие одной или нескольких строк "Single Instance already created-2" означает,
   что для одной или несколько горутин в момент вызова, объект-одиночка уже был
   создан и она не смогла обойти первую проверку.
   
## Другие способы создания объекта-одиночки в Go

* функция **init**

Мы можем создать единственный экземпляр внутри функции `init`. Этот способ 
используется только в том случае, когда возможна ранняя инициализация объекта.
Функция `init` вызывается только один раз для каждого файла в пакете, поэтому
мы можем быть уверены, что будет создан только один экземпляр.

* **sync.Once**

`sync.Once` выполнит операцию единожды. Смотри код ниже:

```go
var once sync.Once

type single struct {
}

var singleInstance *single

func getInstance() *single {
    var once sync.Once
    if singleInstance == nil {
        once.Do(func() {
            fmt.Println("Creating Single Instance Now")
            singleInstance = &single{}
        })
    } else {
        fmt.Println("Single Instance already created-2")
    }
    return singleInstance
}
```

Результат в терминале:

```shell
go run .
Creating Single Instance Now
Single Instance already created-2
Single Instance already created-2
Single Instance already created-2
...
```

* Выводится только одна строка "Creating Single Instance Now", т. е. только
  одна горутина смогла создать экземпляр объекта-одиночки.
* Несколько раз выводится строка "Single Instance already created-2", т. е. для
  нескольких горутин в момент вызова, объект-одиночка уже был создан и они не 
  смогли пройти первую проверку.
