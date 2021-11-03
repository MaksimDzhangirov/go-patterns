package adidas

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/abstract-factory/code/common"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/abstract-factory/code/interfaces"
)

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