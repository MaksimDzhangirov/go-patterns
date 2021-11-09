package main

import "github.com/MaksimDzhangirov/go-concurrency-patterns/patterns/structural/adapter/code/computer"

func main() {
	client := computer.NewClient()
	mac := computer.NewMac()
	client.InsertSquareUsbInComputer(mac)
	windowsMachine := computer.NewWindows()
	windowsMachineAdapter := computer.NewWindowsAdapter(windowsMachine)
	client.InsertSquareUsbInComputer(windowsMachineAdapter)
}