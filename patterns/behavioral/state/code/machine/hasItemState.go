package machine

import (
	"fmt"
)

type hasItemState struct {
	vendingMachine *VendingMachine
}

func NewHasItemState(vendingMachine *VendingMachine) *hasItemState {
	return &hasItemState{
		vendingMachine: vendingMachine,
	}
}

func (i *hasItemState) RequestItem() error {
	if i.vendingMachine.ItemCount == 0 {
		i.vendingMachine.SetState(i.vendingMachine.NoItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requested\n")
	i.vendingMachine.SetState(i.vendingMachine.ItemRequested)
	return nil
}

func (i *hasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.IncrementItemCount(count)
	return nil
}

func (i *hasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *hasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}
