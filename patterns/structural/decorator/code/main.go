package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/decorator/code/pizza"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/decorator/code/topping"
)

func main() {
	veggiePizza := pizza.NewVeggeMania()

	// Добавляем сырную начинку
	veggiePizzaWithCheese := topping.NewCheeseTopping(veggiePizza)

	// Добавляем томатную начинку
	veggiePizzaWithCheeseAndTomato := topping.NewTomatoTopping(veggiePizzaWithCheese)

	fmt.Printf("Price of veggieMania pizza with tomato and cheese topping is %d\n", veggiePizzaWithCheeseAndTomato.GetPrice())

	peppyPaneerPizza := pizza.NewPeppyPaneer()

	// Добавляем сырную начинку
	peppyPaneerPizzaWithCheese := topping.NewCheeseTopping(peppyPaneerPizza)

	fmt.Printf("Price of peppyPaneer with cheese topping is %d\n", peppyPaneerPizzaWithCheese.GetPrice())
}
