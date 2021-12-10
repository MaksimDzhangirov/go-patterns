package machine

import (
	"fmt"
)

type itemRequestedState struct {
	vendingMachine *VendingMachine
}

func NewItemRequestedState(vendingMachine *VendingMachine) *itemRequestedState {
	return &itemRequestedState{
		vendingMachine: vendingMachine,
	}
}

func (i *itemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *itemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *itemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.ItemPrice {
		return fmt.Errorf("insert money is less. Please insert %d", i.vendingMachine.ItemPrice)
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.SetState(i.vendingMachine.HasMoney)
	return nil
}

func (i *itemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}
