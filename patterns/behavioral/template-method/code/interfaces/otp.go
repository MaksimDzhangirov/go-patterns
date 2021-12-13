package interfaces

type IOtp interface {
	GenRandomOTP(int) string
	SaveOTPCache(string)
	GetMessage(string) string
	SendNotification(string) error
	PublishMetric()
}
