package interfaces

type Server interface {
	HandleRequest(string, string) (int, string)
}
