package http

import (
	"github.com/gorilla/mux"
	"github.com/hobord/invst-portfolio-backend-golang/delivery/http/handlers"
	interactor "github.com/hobord/invst-portfolio-backend-golang/interactors"
)

func MakeRouting(router *mux.Router, instrumentInteractor interactor.InstrumentInteractorInterface) {
	instrumentApp := handlers.CreateInstrumentRestHTTPModule(instrumentInteractor)

	router.HandleFunc("/instruments", instrumentApp.Create).Methods("POST")
	router.HandleFunc("/instruments/{id}", instrumentApp.GetByID)
	router.HandleFunc("/instruments", instrumentApp.GetAll).Methods("GET")
	router.HandleFunc("/instruments/{keyword}", instrumentApp.Search).Methods("GET")
	router.HandleFunc("/instruments", instrumentApp.Update).Methods("PUT")
	router.HandleFunc("/instruments/{id}", instrumentApp.Delete).Methods("DELETE")
}
