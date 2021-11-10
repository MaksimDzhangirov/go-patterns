package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/flyweight/code/dress"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/flyweight/code/player"
)

func main() {
	player1 := player.NewPlayer("T", dress.TerroristDressType)
	player2 := player.NewPlayer("T", dress.TerroristDressType)
	player3 := player.NewPlayer("CT", dress.CounterTerroristDressType)
	player4 := player.NewPlayer("CT", dress.CounterTerroristDressType)
	fmt.Printf("Player 1 type %s, dress color %s\n", player1.GetPlayerType(), player1.GetDressColor())
	fmt.Printf("Player 2 type %s, dress color %s\n", player2.GetPlayerType(), player2.GetDressColor())
	fmt.Printf("Player 3 type %s, dress color %s\n", player3.GetPlayerType(), player3.GetDressColor())
	fmt.Printf("Player 4 type %s, dress color %s\n", player4.GetPlayerType(), player4.GetDressColor())
}
