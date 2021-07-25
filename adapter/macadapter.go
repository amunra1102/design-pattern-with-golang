package adapter

type MacAdapter struct {
	M Mac
}

func (m *MacAdapter) insertInSquarePort() {
	m.M.insertInCirclePort()
}
