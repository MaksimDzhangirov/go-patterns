# Шаблон проектирования "Абстрактная фабрика" в Go

[Оригинал](https://golangbyexample.com/abstract-factory-design-pattern-go/)

## Определение

Абстрактная фабрика — это шаблон проектирования, который позволяет вам создавать
семейство связанных объектов. Это абстракция фабричного шаблона. Проще всего
понять этот шаблон на примере. Пусть у нас есть две фабрики

* **nike**
* **adidas**

Представьте, что вам нужно купить комплект спортивной одежды, состоящий из 
**обуви** и **шорт**. В большинстве случаев предпочтительно покупать комплект у одной и
той же фабрики, т. е. nike или adidas. Именно для этого используется абстрактная 
фабрика, таким образом, товары, которые вам нужны (**обувь** и **шорты**) будут 
созданы абстрактной фабрикой **nike** и **adidas**.

Обе эти фабрики **nike** и **adidas** реализуют интерфейс **SportsFactory**.
У нас будет два интерфейса для товаров:
* **Shoe** - этот интерфейс реализован в товарах **nike.shoe** и **adidas.shoe**
* **Short** – этот интерфейс реализован в товарах **nike.short** и **adidas.short**.

Код показан ниже:

**interfaces/sportsFactory.go**

```go
package interfaces

type SportsFactory interface {
    MakeShoe() Shoe
    MakeShort() Short
}
```

**interfaces/shoe.go**

```go
package interfaces

type Shoe interface {
    SetLogo(logo string)
    SetSize(size int)
    GetLogo() string
    GetSize() int
}
```

**interfaces/short.go**

```go
package interfaces

type Short interface {
    SetLogo(log string)
    SetSize(size int)
    GetLogo() string
    GetSize() int
}
```

**common/shoe.go**

```go
package common

type Shoe struct {
    logo string
    size int
}

func NewShoe(logo string, size int) Shoe {
    return Shoe{
        logo: logo,
        size: size,
    }
}

func (s *Shoe) SetLogo(logo string) {
    s.logo = logo
}

func (s *Shoe) GetLogo() string {
    return s.logo
}

func (s *Shoe) SetSize(size int) {
    s.size = size
}

func (s *Shoe) GetSize() int {
    return s.size
}
```

**common/short.go**

```go
package common

type Short struct {
    logo string
    size int
}

func NewShort(logo string, size int) Short {
    return Short{
        logo: logo,
        size: size,
    }
}

func (s *Short) SetLogo(logo string) {
    s.logo = logo
}

func (s *Short) GetLogo() string {
    return s.logo
}

func (s *Short) SetSize(size int) {
    s.size = size
}

func (s *Short) GetSize() int {
    return s.size
}
```

**adidas/shoe.go**

```go
type shoe struct {
    common.Shoe
}
```

**adidas/short.go**

```go
type short struct {
    common.Short
}
```

**adidas/company.go**

```go
type Company struct {
}

func (f *Company) MakeShoe() interfaces.Shoe {
    fmt.Println("Make adidas shoe")
    return &shoe{
        Shoe: common.NewShoe("adidas", 14),
    }
}

func (f *Company) MakeShort() interfaces.Short {
    fmt.Println("Make adidas short")
    return &short{
        Short: common.NewShort("adidas", 14),
    }
}
```

**nike/shoe.go**

```go
type shoe struct {
    common.Shoe
}
```

**nike/short.go**

```go
type short struct {
    common.Short
}
```

**nike/company.go**

```go
type Company struct {
}

func (f *Company) MakeShoe() interfaces.Shoe {
    fmt.Println("Make nike shoe")
    return &shoe{
        Shoe: common.NewShoe("nike", 14),
    }
}

func (f *Company) MakeShort() interfaces.Short {
    fmt.Println("Make nike short")
    return &short{
        Short: common.NewShort("nike", 14),
    }
}
```

**main.go**

```go
func getSportsFactory(brand string) (interfaces.SportsFactory, error) {
    if brand == "adidas" {
        return &adidas.Company{}, nil
    }
    if brand == "nike" {
        return &nike.Company{}, nil
    }
    return nil, fmt.Errorf("wrong brand type passed")
}

func main() {
    adidasFactory, _ := getSportsFactory("adidas")
    nikeFactory, _ := getSportsFactory("nike")
    nikeShoe := nikeFactory.MakeShoe()
    nikeShort := nikeFactory.MakeShort()
    adidasShoe := adidasFactory.MakeShoe()
    adidasShort := adidasFactory.MakeShort()
    printShoeDetails(nikeShoe)
    printShortDetails(nikeShort)
    printShoeDetails(adidasShoe)
    printShortDetails(adidasShort)
}

func printShoeDetails(s interfaces.Shoe) {
    fmt.Printf("Logo: %s\n", s.GetLogo())
    fmt.Printf("Size: %d\n", s.GetSize())
}

func printShortDetails(s interfaces.Short) {
    fmt.Printf("Logo: %s\n", s.GetLogo())
    fmt.Printf("Size: %d\n", s.GetSize())
}
```

Результат в терминале:

```shell
go run main.go 
Make nike shoe
Make nike short
Make adidas shoe
Make adidas short
Logo: nike
Size: 14
Logo: nike
Size: 14
Logo: adidas
Size: 14
Logo: adidas
Size: 14
```
