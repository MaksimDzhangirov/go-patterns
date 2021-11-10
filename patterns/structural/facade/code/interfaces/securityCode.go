package interfaces

type SecurityCode interface {
	CheckCode(int) error
}
