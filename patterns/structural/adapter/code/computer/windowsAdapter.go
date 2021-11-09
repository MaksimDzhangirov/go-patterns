package computer

type windowsAdapter struct {
	windowMachine *windows
}

func NewWindowsAdapter(machine *windows) *windowsAdapter {
	return &windowsAdapter{
		windowMachine: machine,
	}
}

func (w *windowsAdapter) InsertInSquarePort() {
	w.windowMachine.insertInCirclePort()
}