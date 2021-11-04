package common

type House struct {
	windowType string
	doorType   string
	floor      int
}

func NewHouse(windowType string, doorType string, floor int) House {
	return House{
		windowType: windowType,
		doorType:   doorType,
		floor:      floor,
	}
}

func (h *House) GetWindowType() string {
	return h.windowType
}

func (h *House) GetDoorType() string {
	return h.doorType
}

func (h *House) GetFloor() int {
	return h.floor
}
