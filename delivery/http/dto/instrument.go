package dto

type Instrument struct {
	ID     int    `json:"instrumentId"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Type   string `json:"instrumentType"`
}
