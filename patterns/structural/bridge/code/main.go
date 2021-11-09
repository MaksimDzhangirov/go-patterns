package main

import (
	"fmt"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/bridge/code/computer"
	"github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/bridge/code/printer"
)

func main() {
	hpPrinter := printer.NewHp()
	epsonPrinter := printer.NewEpson()
	macComputer := computer.NewMac(hpPrinter)
	macComputer.Print()
	fmt.Println()
	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()
	winComputer := computer.NewWindows(hpPrinter)
	winComputer.Print()
	fmt.Println()
	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
