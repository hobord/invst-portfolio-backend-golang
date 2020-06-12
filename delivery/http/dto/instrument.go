package dto

import "errors"

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

func (instrument *CreateInstrument) Validate() error {
	if instrument.Name == "" {
		return errors.New("missing 'name' property")
	}
	if instrument.Symbol == "" {
		return errors.New("missing 'symbol' property")
	}
	if instrument.Type == "" {
		return errors.New("missing 'type' property")
	}
	return nil
}

type ListOfInstruments struct {
	Data  []Instrument `json:"data"`
	Total int          `json:"total"`
}

type DeleteIntrumentResponse struct {
	Success bool `json:"success"`
}
