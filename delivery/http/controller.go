package http

import (
	"github.com/gorilla/mux"
	"github.com/hobord/invst-portfolio-backend-golang/delivery/http/handlers"
	interactor "github.com/hobord/invst-portfolio-backend-golang/interactors"
)

func MakeRouting(router *mux.Router, spaDir string, instrumentInteractor interactor.InstrumentInteractorInterface) {
	instrumentApp := handlers.CreateInstrumentRestHTTPModule(instrumentInteractor)
	spaHandler := handlers.SpaHandler(spaDir, "index.html")

	router.HandleFunc("/instruments", instrumentApp.CreateInstrument).Methods("POST")
	router.HandleFunc("/instruments/{id:[0-9]+}", instrumentApp.GetInstrumentByID).Methods("GET")
	router.HandleFunc("/instruments", instrumentApp.ListInstrument).Methods("GET")
	router.HandleFunc("/instruments", instrumentApp.UpdateInstrument).Methods("PUT")
	router.HandleFunc("/instruments/{id:[0-9]+}", instrumentApp.DeleteInstrument).Methods("DELETE")
	if spaDir != "" {
		router.PathPrefix("/").Handler(spaHandler)
	}
}
