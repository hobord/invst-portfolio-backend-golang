package http

import (
	"github.com/gorilla/mux"
	interactor "github.com/hobord/invst-portfolio-backend-golang/interactors"
)

func MakeRouting(router *mux.Router, instrumentInteractor interactor.InstrumentInteractorInterface) {
	// httpApp := handlers.CreateRestHTTPModule(instrumentInteractor)
}
