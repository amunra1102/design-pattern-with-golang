package builder

type House struct {
	windowType string
	doorType string
	floor int
}

// GetWindowType is getter
func (h *House) GetWindowType() string {
	return h.windowType
}

// GetDoorType is getter
func (h *House) GetDoorType() string {
	return h.doorType
}

// GetFloor is getter
func (h *House) GetFloor() int {
	return h.floor
}
