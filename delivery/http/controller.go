package http

import (
	"github.com/gorilla/mux"
	"github.com/hobord/invst-portfolio-backend-golang/delivery/http/handlers"
	interactor "github.com/hobord/invst-portfolio-backend-golang/interactors"
)

func MakeRouting(router *mux.Router, instrumentInteractor interactor.InstrumentInteractorInterface) {
	instrumentApp := handlers.CreateInstrumentRestHTTPModule(instrumentInteractor)

	router.HandleFunc("/instruments", instrumentApp.CreateInstrument).Methods("POST")
	router.HandleFunc("/instruments/{id}", instrumentApp.GetInstrumentByID)
	// router.HandleFunc("/instruments", instrumentApp.GetAllInstrument).Methods("GET")
	router.HandleFunc("/instruments", instrumentApp.ListInstrument).Methods("GET")

	router.HandleFunc("/instruments", instrumentApp.UpdateInstrument).Methods("PUT")
	router.HandleFunc("/instruments/{id}", instrumentApp.DeleteInstrument).Methods("DELETE")
}
