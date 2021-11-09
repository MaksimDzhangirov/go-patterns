package topping

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/decorator/code/interfaces"

type tomatoTopping struct {
	pizza interfaces.Pizza
}

func NewTomatoTopping(pizza interfaces.Pizza) *tomatoTopping {
	return &tomatoTopping{pizza: pizza}
}

func (t *tomatoTopping) GetPrice() int {
	pizzaPrice := t.pizza.GetPrice()
	return pizzaPrice + 7
}
