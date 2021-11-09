package topping

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/decorator/code/interfaces"

type cheeseTopping struct {
	pizza interfaces.Pizza
}

func NewCheeseTopping(pizza interfaces.Pizza) *cheeseTopping {
	return &cheeseTopping{pizza: pizza}
}

func (t *cheeseTopping) GetPrice() int {
	pizzaPrice := t.pizza.GetPrice()
	return pizzaPrice + 10
}
