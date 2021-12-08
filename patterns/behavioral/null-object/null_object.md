# Шаблон проектирования "Объект Null" в Go

[Оригинал](https://golangbyexample.com/null-object-design-pattern-golang/)

## Введение

Шаблон "Объект Null" - это поведенческий шаблон проектирования. Он полезен,
когда клиентский код полагается на некую зависимость, которая может быть 
равна Null. Использование этого шаблона проектирования избавляет клиентов от
необходимости выполнять проверки на Null результатов этих зависимостей. С 
учетом сказанного следует также отметить, что клиент ведёт себя нормально, когда
зависимости равны Null.

Основными компонентами шаблона проектирования "Объект Null" являются:

1. **Entity** – интерфейс, определяющий примитивные операции, которые дочерние 
   структуры должны реализовать.
2. **ConcreteEntity** - реализует интерфейс **Entity**.
3. **NullEntity** - представляет "Объект Null". Он также реализует интерфейс 
   сущности, но имеет null значения.
4. **Client** - клиент получает реализацию интерфейса сущности и использует её. 
   Ему всё равно какая реализация пришла на вход: **ConcreteEntity** или 
   **NullEntity**. Он может работать с любой из них.
   
Рассмотрим пример. Предположим, у нас есть колледж с множеством факультетов, на
каждом из которых работает несколько профессоров.

Факультет (**department**) представим с помощью интерфейса

```go
type department interface {
    getNumberOfProfessors() int
    getName() string
}
```

тогда как колледж следующим образом:

```go
type college struct {
    departments []department
}
```

Теперь предположим, что существует агентство, которое хочет подсчитать общее 
количество профессоров в конкретном колледже только для определенных факультетов.
Мы будем использовать здесь шаблон проектирования "Объект Null". Колледж будет 
возвращать объект `nullDepartment` (смотри **nullDepartment.go**), если такого
факультета нет в колледже.

Обратите внимание на код в **agency.go**

* **agency.go** даже не волнует существует ли в колледже конкретный факультет или 
  нет. `college` возвращает `nullDepartment`, если такого факультета нет в колледже.
* работа с `nullDepartment` и реально существующим факультетом ничем не отличаются,
  поэтому проверки на null не нужны. **agency.go** вызывает метод 
  `getNumberOfProfessors()` для обоих объектов.

Выше приведены два преимущества, которые мы получаем, используя шаблон 
проектирования "Объект Null" для этого случая. Код показан ниже

**agency.go**

```go
func main() {
    c1 := createCollege1()
    c2 := createCollege2()
    totalProfessors := 0
    departmentArray := []string{"computerScience", "mechanical", "civil", "electronics"}
    
    for _, departmentName := range departmentArray {
        d := c1.GetDepartment(departmentName)
        totalProfessors += d.GetNumberOfProfessors()
    }
    fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)
    
    // Обнуляем общее число профессоров
    totalProfessors = 0
    for _, departmentName := range departmentArray {
        d := c2.GetDepartment(departmentName)
        totalProfessors += d.GetNumberOfProfessors()
    }
    fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}

func createCollege1() *college.College {
    c := college.NewCollege()
    c.AddDepartment("computerScience", 4)
    c.AddDepartment("mechanical", 5)
    return c
}

func createCollege2() *college.College {
    c := college.NewCollege()
    c.AddDepartment("computerScience", 2)
    return c
}
```

**college/college.go** - описывает колледж

```go
type College struct {
    departments []interfaces.Department
}

func NewCollege() *College {
    return &College{}
}

func (c *College) AddDepartment(departmentName string, numOfProfessors int) {
    if departmentName == "computerScience" {
        computerScienceDepartment := department.NewComputerScience(numOfProfessors)
        c.departments = append(c.departments, computerScienceDepartment)
    }
    if departmentName == "mechanical" {
        mechanicalDepartment := department.NewMechanical(numOfProfessors)
        c.departments = append(c.departments, mechanicalDepartment)
    }
    return
}

func (c *College) GetDepartment(departmentName string) interfaces.Department {
    for _, currentDepartment := range c.departments {
        if currentDepartment.GetName() == departmentName {
            return currentDepartment
        }
    }
    // Возвращаем nullDepartment, если такого факультета не существует
    return department.NewNullDepartment()
}
```

**interfaces/department.go** - описывает интерфейс `department`

```go
type Department interface {
    GetNumberOfProfessors() int
    GetName() string
}
```

**college/department/computerScience.go** – конкретная реализация интерфейса `department`

```go
type computerScience struct {
    numberOfProfessors int
}

func NewComputerScience(numberOfProfessors int) *computerScience {
    return &computerScience{
        numberOfProfessors: numberOfProfessors,
    }
}

func (c *computerScience) GetNumberOfProfessors() int {
    return c.numberOfProfessors
}

func (c *computerScience) GetName() string {
    return "computerScience"
}
```

**college/department/mechanical.go** – конкретная реализация интерфейса `department`

```go
type mechanical struct {
    numberOfProfessors int
}

func NewMechanical(numberOfProfessors int) *mechanical {
    return &mechanical{
        numberOfProfessors: numberOfProfessors,
    }
}

func (m *mechanical) GetNumberOfProfessors() int {
    return m.numberOfProfessors
}

func (m *mechanical) GetName() string {
    return "mechanical"
}
```

**college/department/nullDepartment.go** – реализация "Объекта Null" интерфейса `department`

```go
type nullDepartment struct {
    numberOfProfessors int
}

func NewNullDepartment() *nullDepartment {
    return &nullDepartment{}
}

func (d *nullDepartment) GetNumberOfProfessors() int {
    return 0
}

func (d *nullDepartment) GetName() string {
    return "nullDepartment"
}
```

Результат в терминале:

```shell
go run agency.go
Total number of professors in college1 is 9
Total number of professors in college2 is 2
```