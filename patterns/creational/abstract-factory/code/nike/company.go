package nike

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/abstract-factory/code/common"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/creational/abstract-factory/code/interfaces"
)

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