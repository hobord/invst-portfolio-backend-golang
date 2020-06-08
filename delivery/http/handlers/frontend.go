package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

// Serve from a public directory with specific index
type spaHandler struct {
	publicDir string // The directory from which to serve
	indexFile string // The fallback/default file to serve
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := filepath.Join(h.publicDir, filepath.Clean(r.URL.Path))

	if info, err := os.Stat(p); err != nil {
		http.ServeFile(w, r, filepath.Join(h.publicDir, h.indexFile))
		return
	} else if info.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.publicDir, h.indexFile))
		return
	}

	http.ServeFile(w, r, p)
}

// Returns a request handler (http.Handler) that serves a single
// page application from a given public directory (publicDir).
func SpaHandler(publicDir string, indexFile string) http.Handler {
	return &spaHandler{publicDir, indexFile}
}
