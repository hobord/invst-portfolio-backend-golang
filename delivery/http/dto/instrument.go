package dto

type Instrument struct {
	ID     int    `json:"instrumentId"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Type   string `json:"instrumentType"`
}

type CreateInstrument struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Type   string `json:"instrumentType"`
}

type ListOfInstruments struct {
	Data  []Instrument `json:"data"`
	Total int          `json:"total"`
}
