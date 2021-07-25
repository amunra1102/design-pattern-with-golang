package flyweight

type Factory struct {
	soldierMap map[string]ISoldier
}

func NewFactory() *Factory {
	return &Factory{soldierMap: make(map[string]ISoldier)}
}

func (f *Factory) CreateSoldier(name string) ISoldier {
	if val, ok := f.soldierMap[name]; ok {
		return val
	}

	f.soldierMap[name] = newSoldier(name)
	return f.soldierMap[name]
}

func (f *Factory) GetSize() int {
	return len(f.soldierMap)
}
