package entity

type Instrument struct {
	ID     int
	Name   string
	Symbol string
	Type   string
}

func CreateInstrumentEntity(name, symbol, instumentType string) *Instrument {
	return &Instrument{
		ID:     0,
		Name:   name,
		Symbol: symbol,
		Type:   instumentType,
	}
}
