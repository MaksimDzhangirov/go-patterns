package account

import "fmt"

type account struct {
	name string
}

func NewAccount(accountName string) *account {
	return &account{
		name: accountName,
	}
}

func (a *account) CheckAccount(accountName string) error {
	if a.name != accountName {
		return fmt.Errorf("account name is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
