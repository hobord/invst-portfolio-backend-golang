package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hobord/invst-portfolio-backend-golang/delivery/http/dto"
	"github.com/hobord/invst-portfolio-backend-golang/domain/entity"
	interactor "github.com/hobord/invst-portfolio-backend-golang/interactors"
)

// InstrumentRestHTTPModule is handle the entity related http request responses
type InstrumentRestHTTPModule struct {
	instrumentInteractor interactor.InstrumentInteractorInterface
}

// CreateInstrumentRestHTTPModule create a new http handler app to entity
func CreateInstrumentRestHTTPModule(instrumentInteractor interactor.InstrumentInteractorInterface) *InstrumentRestHTTPModule {
	return &InstrumentRestHTTPModule{
		instrumentInteractor: instrumentInteractor,
	}
}

func makeInstrumentDTO(instrument *entity.Instrument) *dto.Instrument {
	return &dto.Instrument{
		ID:     instrument.ID,
		Name:   instrument.Name,
		Symbol: instrument.Symbol,
		Type:   instrument.Type,
	}
}

// GetByID return entity by id
func (app *InstrumentRestHTTPModule) GetInstrumentByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id should be integer", http.StatusBadRequest)
		return
	}

	entity, err := app.instrumentInteractor.GetByID(r.Context(), i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		err = errors.New("no resource found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	entityDTO := makeInstrumentDTO(entity)

	js, err := json.Marshal(entityDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// List return list of entities
func (app *InstrumentRestHTTPModule) ListInstrument(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("key")

	res, _, err := app.instrumentInteractor.List(r.Context(), keyword, 0, 100)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entityDTOs := make([]*dto.Instrument, 0)
	for _, entity := range res {
		entityDTO := makeInstrumentDTO(entity)
		entityDTOs = append(entityDTOs, entityDTO)
	}

	// Convert to json
	js, err := json.Marshal(entityDTOs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back to response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Create is update to persistent the entity
func (app *InstrumentRestHTTPModule) CreateInstrument(w http.ResponseWriter, r *http.Request) {
	// Decode the request DTO.
	decoder := json.NewDecoder(r.Body)
	var createDTO dto.CreateInstrument
	err := decoder.Decode(&createDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new entity.
	entity := entity.CreateInstrumentEntity(createDTO.Name, createDTO.Symbol, createDTO.Type)

	// Save the new entity.
	err = app.instrumentInteractor.Save(r.Context(), entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new response DTO.
	entityDTO := makeInstrumentDTO(entity)

	// Convert to json
	js, err := json.Marshal(entityDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back to response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Update is update to persistent the entity.
func (app *InstrumentRestHTTPModule) UpdateInstrument(w http.ResponseWriter, r *http.Request) {
	// Decode the request DTO.
	decoder := json.NewDecoder(r.Body)
	var updateDTO dto.Instrument
	err := decoder.Decode(&updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if updateDTO.ID == 0 {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	// Load the original entity.
	entity, err := app.instrumentInteractor.GetByID(r.Context(), updateDTO.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		err = errors.New("no resource found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Update the entity properties.
	entity.Name = updateDTO.Name
	entity.Symbol = updateDTO.Symbol
	entity.Type = updateDTO.Type

	// save the entity
	err = app.instrumentInteractor.Save(r.Context(), entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a response DTO.
	entityDTO := makeInstrumentDTO(entity)

	// Convert to json
	js, err := json.Marshal(entityDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back to response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Delete entity
func (app *InstrumentRestHTTPModule) DeleteInstrument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}
	i, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "id should be integer", http.StatusBadRequest)
		return
	}

	entity, err := app.instrumentInteractor.GetByID(r.Context(), i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = app.instrumentInteractor.Delete(r.Context(), entity.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
