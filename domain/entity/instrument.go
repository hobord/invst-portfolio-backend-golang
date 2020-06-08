package entity

type Instrument struct {
	ID     int
	Name   string
	Symbol string
	Type   string
}

func CreateInstrumentEntity(name, symbol, _type string) *Instrument {
	return &Instrument{
		ID:     0,
		Name:   name,
		Symbol: symbol,
		Type:   _type,
	}
}
