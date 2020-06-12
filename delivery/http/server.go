package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	customHandlers "github.com/hobord/invst-portfolio-backend-golang/delivery/http/handlers"
	interactor "github.com/hobord/invst-portfolio-backend-golang/interactors"
)

func MakeWebServer(httpPort int, allowedOrigins []string, spaDir string, interactor *interactor.InstrumentInteractor) {
	r := mux.NewRouter()
	MakeRouting(r, spaDir, interactor)

	methods := handlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT", "PATCH"})
	origins := handlers.AllowedOrigins(allowedOrigins)
	handler := handlers.CORS(methods, origins)(r)
	logHandler := customHandlers.LogRequestHandler(handler)

	log.Printf("Listen on port: %v", httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logHandler))
}
