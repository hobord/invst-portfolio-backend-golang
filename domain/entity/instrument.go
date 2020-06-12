package entity

type Instrument struct {
	ID     int
	Name   string
	Symbol string
	Type   string
}

func CreateInstrumentEntity(id int, name, symbol, instumentType string) *Instrument {
	return &Instrument{
		ID:     id,
		Name:   name,
		Symbol: symbol,
		Type:   instumentType,
	}
}
