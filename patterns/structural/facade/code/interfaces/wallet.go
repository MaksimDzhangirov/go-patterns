package interfaces

type Wallet interface {
	CreditBalance(int)
	DebitBalance(int) error
}
