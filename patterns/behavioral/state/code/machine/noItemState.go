package machine

import (
	"fmt"
)

type noItemState struct {
	vendingMachine *VendingMachine
}

func NewNoItemState(vendingMachine *VendingMachine) *noItemState {
	return &noItemState{
		vendingMachine: vendingMachine,
	}
}

func (i *noItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) AddItem(count int) error {
	i.vendingMachine.IncrementItemCount(count)
	i.vendingMachine.SetState(i.vendingMachine.HasItem)
	return nil
}

func (i *noItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *noItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}
