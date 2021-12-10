package machine

import (
	"fmt"
)

type hasMoneyState struct {
	vendingMachine *VendingMachine
}

func NewHasMoneyState(vendingMachine *VendingMachine) *hasMoneyState {
	return &hasMoneyState{
		vendingMachine: vendingMachine,
	}
}

func (i *hasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *hasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *hasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.ItemCount = i.vendingMachine.ItemCount - 1
	if i.vendingMachine.ItemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.NoItem)
	} else {
		i.vendingMachine.SetState(i.vendingMachine.HasItem)
	}
	return nil
}
