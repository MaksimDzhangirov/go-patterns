package interfaces

type Notification interface {
	SendWalletCreditNotification()
	SendWalletDebitNotification()
}
